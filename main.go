package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/thalescpl-io/ms-dke-api/dke/v1"
	"github.com/thalescpl-io/ms-dke-api/keystore"
)

type Server struct {
	keyStore keystore.KeyStore
}

func newServer(keyStore keystore.KeyStore) dke.AzureDKEKeyManagementServiceServer {
	return &Server{
		keyStore: keyStore,
	}
}

func (s *Server) GetPublicKey(ctx context.Context, req *dke.PublicKeyRequest) (*dke.PublicKeyResponse, error) {
	key, err := s.keyStore.GetActiveKey(req.KeyName)
	if err != nil {
		return nil, err
	}

	pubKey := key.PublicKey().(rsa.PublicKey)

	return &dke.PublicKeyResponse{
		Key: &dke.PublicKeyResponse_Key{
			Kty: key.Type(),
			N:   pubKey.N.Bytes(),
			E:   int32(pubKey.E),
			Alg: key.Algorithm(),
			Kid: fmt.Sprintf("%s/%s", req.KeyName, key.Id()),
		},
	}, nil
}

func (s *Server) Decrypt(ctx context.Context, req *dke.DecryptRequest) (*dke.DecryptResponse, error) {
	key, err := s.keyStore.GetKey(req.KeyName, req.KeyId)
	if err != nil {
		return nil, err
	}

	decryptedBytes, err := key.Decrypt(req.Value, req.Alg)
	if err != nil {
		return nil, err
	}

	return &dke.DecryptResponse{
		Value: decryptedBytes,
	}, nil
}

func main() {
	flag.Parse()

	keyPem, err := ioutil.ReadFile("key.pem")
	if err != nil {
		panic(err)
	}

	keyBlock, _ := pem.Decode([]byte(keyPem))
	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	keyStore := keystore.NewPlain()
	keyStore.AddRSAKey("TestKey1", "D798B899-3350-4F5C-A608-2EDA37CB0EBD", key)

	var opts []runtime.ServeMuxOption
	server := runtime.NewServeMux(opts...)
	dke.RegisterAzureDKEKeyManagementServiceHandlerServer(context.Background(), server, newServer(keyStore))

	addr := "localhost:5000"
	fmt.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, http.HandlerFunc(server.ServeHTTP)); err != nil {
		panic(err)
	}
}
