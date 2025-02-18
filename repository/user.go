package repository

import (
	"fmt"
	"gofr-test/model"

	"gofr.dev/pkg/gofr"
)

type UserRepository struct {
}

func newUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUsers(ctx *gofr.Context) (any, error) {
	var users []model.User
	rows, err := ctx.SQL.QueryContext(ctx, "SELECT * FROM test")

	fmt.Println(err, ".......................")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// return the customer
	return users, nil

}
