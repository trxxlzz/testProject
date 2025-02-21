package service

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string, password string) (int64, string, error)
}

type UserService struct {
	Repo UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, name string, password string) (id int64, message string, err error) {
	id, message, err = s.Repo.CreateUser(ctx, name, password)
	if err != nil {
		return 0, "", err
	}

	return id, "User succesfully created!", err
}
