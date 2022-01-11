package storage

import (
	"amarbank/module/user_v1/domain"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
)

func (s *userFileStorage) Create(ctx context.Context, input *domain.UserModel) error {

	fd, _ := os.Open(s.path)

	//TODO: make sure file is closed after read and write done
	defer fd.Close()

	var user []domain.UserModel

	fileData, err := ioutil.ReadAll(fd)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileData, &user); err != nil {
		return err
	}

	user = append(user, *input)

	userBytes, err := json.Marshal(user)
	fd.Name()
	if err := ioutil.WriteFile(fd.Name(), userBytes, 0664); err != nil {
		return err
	}

	return nil
}
