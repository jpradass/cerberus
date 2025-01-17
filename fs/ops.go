package fs

import (
	"fmt"
	"os"
)

func IsPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func ReadContent(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error: there was an error reading content: %w", err)
	}

	return content, nil
}
