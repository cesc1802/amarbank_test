package storage

import (
	"amarbank/module/user_v1/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

func (s *userFileStorage) GetByID(ctx context.Context, uuid uuid.UUID) (*domain.UserModel, error) {
	fd, _ := os.Open(s.path)

	var users []domain.UserModel

	userBytesData, _ := ioutil.ReadAll(fd)

	json.Unmarshal(userBytesData, &users)

	for idx := range users {
		if users[idx].ID.String() == uuid.String() {
			return &users[idx], nil
		}
	}

	return nil, fmt.Errorf("user not found")
}
