package encryption

import (
	"testing"

	"golang.org/x/crypto/chacha20poly1305"
)

func TestNewKey(t *testing.T) {
	key, err := NewKey()
	if err != nil {
		t.Error(err)
	}

	if len(key) != chacha20poly1305.KeySize {
		t.Errorf("key length is not what chacha20 algorithm requires. key: %d, wanted: %d", len(key), chacha20poly1305.KeySize)
	}
}

func TestNewNonce(t *testing.T) {
	nonce, err := NewNonce()
	if err != nil {
		t.Error(err)
	}

	if len(nonce) != chacha20poly1305.NonceSize {
		t.Errorf("nonce length is not what chacha20 algorithm requires. nonce: %d, wanted: %d", len(nonce), chacha20poly1305.NonceSize)
	}
}
