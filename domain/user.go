package domain

import "context"

type User struct {
	ID       int64
	Provider string
	UID      string
}

type UserUsecase interface {
	GetByProviderWithUID(ctx context.Context, provider string, uid string) (*User, error)
	Store(ctx context.Context, user *User) error
}

type UserRepository interface {
	GetByProviderWithUID(ctx context.Context, provider string, uid string) (*User, error)
	Store(ctx context.Context, user *User) error
}
