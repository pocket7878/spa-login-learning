package domain

import "context"

type User struct {
	ID    int64
	Email string
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	Store(ctx context.Context, user *User) error
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	Store(ctx context.Context, user *User) error
}
