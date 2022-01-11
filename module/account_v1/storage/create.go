package storage

import (
	"amarbank/module/account_v1/domain"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
)

func (s *accountFileStorage) Create(ctx context.Context, input *domain.AccountModel) error {

	fd, _ := os.Open(s.path)

	//TODO: make sure file is closed after read and write done
	defer fd.Close()

	var accounts []domain.AccountModel

	fileData, err := ioutil.ReadAll(fd)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileData, &accounts); err != nil {
		return err
	}

	accounts = append(accounts, *input)

	userBytes, err := json.Marshal(accounts)
	fd.Name()
	if err := ioutil.WriteFile(fd.Name(), userBytes, 0664); err != nil {
		return err
	}

	return nil
}
