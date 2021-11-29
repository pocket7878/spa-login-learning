package usecase

import (
	"context"

	"github.com/pocket7878/spa_login_learning_backend/domain"
)

type UserUsecaseImpl struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo: repo,
	}
}

func (u *UserUsecaseImpl) GetByProviderWithUID(ctx context.Context, provider string, uid string) (*domain.User, error) {
	return u.repo.GetByProviderWithUID(ctx, provider, uid)
}

func (u *UserUsecaseImpl) Store(ctx context.Context, user *domain.User) error {
	return u.repo.Store(ctx, user)
}
