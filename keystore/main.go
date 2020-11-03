package keystore

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type Key interface {
	Id() string
	Type() string
	Algorithm() string
	PublicKey() interface{}
	Decrypt(encrypted []byte, algo string) ([]byte, error)
}

type KeyStore interface {
	GetActiveKey(keyName string) (Key, error)
	GetKey(keyName, keyId string) (Key, error)
}

type RSAKey struct {
	id      string
	private *rsa.PrivateKey
}

func (k *RSAKey) Id() string {
	return k.id
}

func (k *RSAKey) Type() string {
	return "RSA"
}

func (k *RSAKey) Algorithm() string {
	return "RS256"
}

func (k *RSAKey) PublicKey() interface{} {
	return k.private.PublicKey
}

func (k *RSAKey) Decrypt(encrypted []byte, algo string) ([]byte, error) {
	switch algo {
	case "RSA-OAEP-256":
		return rsa.DecryptOAEP(sha256.New(), rand.Reader, k.private, encrypted, []byte{})
	default:
		return nil, fmt.Errorf("Unsupported algorithm %s", algo)
	}
}

type PlainKeyStore struct {
	nameToKeys map[string]Key
	idToKeys   map[string]Key
}

func NewPlain() *PlainKeyStore {
	return &PlainKeyStore{
		nameToKeys: make(map[string]Key, 1),
		idToKeys:   make(map[string]Key, 1),
	}
}

func (ks *PlainKeyStore) AddRSAKey(keyName, keyId string, privKey *rsa.PrivateKey) error {
	key := RSAKey{
		id:      keyId,
		private: privKey,
	}

	ks.nameToKeys[keyName] = &key
	ks.idToKeys[keyId] = &key

	return nil
}

func (ks *PlainKeyStore) GetActiveKey(keyName string) (Key, error) {
	key, ok := ks.nameToKeys[keyName]
	if !ok {
		return nil, fmt.Errorf("Key not found")
	}
	return key, nil
}

func (ks *PlainKeyStore) GetKey(keyName, keyId string) (Key, error) {
	key, ok := ks.idToKeys[keyId]
	if !ok {
		return nil, fmt.Errorf("Key not found")
	}
	return key, nil
}
