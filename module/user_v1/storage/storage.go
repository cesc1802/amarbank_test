package storage

import (
	"fmt"
	"io/fs"
	"os"
)

type userFileStorage struct {
	path string
	perm fs.FileMode
}

func NewUSerFileStorage(filename string, perm fs.FileMode) (*userFileStorage, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("cannot open file %v", err.Error())
	}

	return &userFileStorage{
		path: fmt.Sprintf("%s/%s", dir, filename),
		perm: perm,
	}, nil
}
