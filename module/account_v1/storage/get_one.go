package storage

import (
	"amarbank/module/account_v1/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

func (s *accountFileStorage) GetByID(ctx context.Context, uuid uuid.UUID) (*domain.AccountModel, error) {
	fd, _ := os.Open(s.path)

	var accounts []domain.AccountModel

	userBytesData, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(userBytesData, &accounts); err != nil {
		return nil, err
	}

	for idx := range accounts {
		if accounts[idx].ID.String() == uuid.String() {
			return &accounts[idx], nil
		}
	}

	return nil, fmt.Errorf("account not found")
}

func (s *accountFileStorage) GetByUserID(ctx context.Context, userUuid uuid.UUID) (*domain.AccountModel, error) {
	fd, _ := os.Open(s.path)

	var accounts []domain.AccountModel

	userBytesData, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(userBytesData, &accounts); err != nil {
		return nil, err
	}

	for idx := range accounts {
		if accounts[idx].UserID.String() == userUuid.String() {
			return &accounts[idx], nil
		}
	}

	return nil, fmt.Errorf("account not found")
}
