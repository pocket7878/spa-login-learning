package usecase

import (
	"context"

	"github.com/pocket7878/spa_login_learning_backend/domain"
)

type UserUsecaseImpl struct {
	repo domain.UserRepository
}

func (u *UserUsecaseImpl) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UserUsecaseImpl) Store(ctx context.Context, user *domain.User) error {
	return u.repo.Store(ctx, user)
}
