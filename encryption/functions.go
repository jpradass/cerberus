package encryption

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

// Creates a new random key as chacha20poly1305 algorithm requires
func NewKey() ([]byte, error) {
	// Generate a random 256-bit (32-byte) key
	key := make([]byte, chacha20poly1305.KeySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed generating key: %w", err)
	}

	return key, nil
}

// Creates a new random nonce as chacha20poly1305 algorithm requires
func NewNonce() ([]byte, error) {
	nonce := make([]byte, chacha20poly1305.NonceSize)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("failed generating nonce: %w", err)
	}

	return nonce, nil
}

func Seal(key, nonce, text []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, fmt.Errorf("There was an error creating AEAD: %w", err)
	}

	return aead.Seal(nil, nonce, text, nil), nil
}

func Unveil(key, nonce, cipher []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, fmt.Errorf("There was an error creating AEAD: %w", err)
	}

	return aead.Open(nil, nonce, cipher, nil)
}
