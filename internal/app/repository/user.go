package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nurhusni/go-graphql/internal/app/entity"
)

type UserRepository interface {
	GetUsers(ctx context.Context, params entity.User) (result []entity.User, err error)
	CreateUser(ctx context.Context, params entity.User) (err error)
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return UserRepositoryImpl{
		DB: DB,
	}
}

func (u UserRepositoryImpl) GetUsers(ctx context.Context, params entity.User) (result []entity.User, err error) {
	result = []entity.User{}

	cond := ""
	if params.Name != "" {
		cond += fmt.Sprintf(" AND name = '%s'", params.Name)
	}

	if params.Email != "" {
		cond += fmt.Sprintf(" AND email = '%s'", params.Email)
	}

	if params.PhoneNumber != "" {
		cond += fmt.Sprintf(" AND phone_number = '%s'", params.PhoneNumber)
	}

	query := fmt.Sprintf(`SELECT name, email, phone_number FROM users WHERE true %s`, cond)
	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		data := entity.User{}
		err = rows.Scan(
			&data.Name,
			&data.Email,
			&data.PhoneNumber,
		)
		if err != nil {
			return
		}

		result = append(result, data)
	}

	return
}

func (u UserRepositoryImpl) CreateUser(ctx context.Context, params entity.User) (err error) {
	_, err = u.DB.ExecContext(ctx, "INSERT INTO users(name, email, phone_number) VALUES($1, $2, $3)",
		params.Name, params.Email, params.PhoneNumber,
	)
	if err != nil {
		return
	}
	return
}
