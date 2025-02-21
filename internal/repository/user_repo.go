package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) CreateUser(ctx context.Context, name string, password string) (int64, string, error) {
	var id int64

	err := r.DB.QueryRowContext(ctx, `INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id`, name, password).Scan(&id)
	if err != nil {
		return 0, "", fmt.Errorf("failed to create user: %w", err)
	}

	return id, "User successfully created!", nil
}
