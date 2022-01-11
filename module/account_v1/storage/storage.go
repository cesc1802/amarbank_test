package storage

import (
	"fmt"
	"os"
)

type accountFileStorage struct {
	path string
	perm os.FileMode
}

func NewAccountFileStorage(fileName string, perm os.FileMode) (*accountFileStorage, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("cannot create account storage %v", err)
	}
	return &accountFileStorage{
		path: fmt.Sprintf("%s/%s", dir, fileName),
		perm: perm,
	}, nil
}
