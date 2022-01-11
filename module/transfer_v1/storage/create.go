package storage

import (
	"amarbank/module/transfer_v1/domain"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
)

func (s *txnFileStorage) Create(ctx context.Context, input *domain.TransactionModel) error {

	fd, _ := os.Open(s.path)

	//TODO: make sure file is closed after read and write done
	defer fd.Close()

	var transactions []domain.TransactionModel

	fileData, err := ioutil.ReadAll(fd)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileData, &transactions); err != nil {
		return err
	}

	transactions = append(transactions, *input)

	userBytes, err := json.Marshal(transactions)
	if err := ioutil.WriteFile(fd.Name(), userBytes, 0664); err != nil {
		return err
	}

	return nil
}
