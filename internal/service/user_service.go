package service

import (
	"context"
	"testProject/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func (s *UserService) CreateUser(ctx context.Context, name string, password string) (id int64, err error) {
	id, err = s.Repo.CreateUser(ctx, name, password)
	if err != nil {
		return 0, err
	}

	return id, err
}
