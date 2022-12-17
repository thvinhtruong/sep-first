package db

import (
	"context"
	"server/UserService/pkg/hasher"
)

type RegisterUserParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type RegisterUserResult struct {
	ID       int64
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
}

func (store *TxStore) RegisterUser(ctx context.Context, arg RegisterUserParams) (RegisterUserResult, error) {
	var result RegisterUserResult
	err := store.enableTx(ctx, func(q *Queries) error {
		user, err := q.CreateUser(ctx, CreateUserParams{
			Fullname: arg.Fullname,
			Phone:    arg.Phone,
			Gender:   arg.Gender,
		})

		if err != nil {
			return err
		}

		id, err := user.LastInsertId()
		if err != nil {
			return err
		}

		hashed, err := hasher.HashPassword(arg.Password)
		if err != nil {
			return err
		}

		err = q.CreateUserPassword(ctx, CreateUserPasswordParams{
			UserID:   int32(id),
			Password: hashed,
		})

		if err != nil {
			return err
		}

		result.ID = id
		result.FullName = arg.Fullname
		result.Phone = arg.Phone
		result.Gender = arg.Gender

		return nil
	})
	return result, err
}

func (store *TxStore) DeleteUser(ctx context.Context, id int32) error {
	return store.enableTx(ctx, func(q *Queries) error {
		err := q.DeleteUser(ctx, id)
		if err != nil {
			return err
		}

		err = q.DeleteUserPassword(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}
