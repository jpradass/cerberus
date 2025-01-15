package main

import (
	"fmt"
	"os"

	"github.com/jpradass/cerberus/encryption"
)

func main() {
	key, err := encryption.NewKey()
	if err != nil {
		fmt.Errorf("There was an error generating key: %w", err)
		os.Exit(1)
	}

	nonce, err := encryption.NewNonce()
	if err != nil {
		fmt.Errorf("There was an error generating nonce: %w", err)
		os.Exit(1)
	}

	plaintext := []byte("Hello, ChaCha20-Poly1305 encryption in Go!")
	encrypted, err := encryption.Seal(key, nonce, plaintext)
	fmt.Printf("encrypted secret: %x\n", encrypted)

	decrypted, err := encryption.Unveil(key, nonce, encrypted)
	fmt.Printf("decrypted secret: %s\n", decrypted)
	// cmd.Execute()
}
