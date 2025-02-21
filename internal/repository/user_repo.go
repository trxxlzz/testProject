package repository

import (
	"context"
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) CreateUser(ctx context.Context, name string, password string) (id int64, err error) {
	result, err := r.DB.ExecContext(ctx, `INSERT INTO users (name, password) VALUES (?, ?)`, name, password)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}
