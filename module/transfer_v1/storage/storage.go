package storage

import (
	"fmt"
	"os"
)

type txnFileStorage struct {
	path string
	perm os.FileMode
}

func NewTxnFileStorage(fileName string, perm os.FileMode) (*txnFileStorage, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("cannot create transaction storage %v", err)
	}
	return &txnFileStorage{
		path: fmt.Sprintf("%s/%s", dir, fileName),
		perm: perm,
	}, nil
}
