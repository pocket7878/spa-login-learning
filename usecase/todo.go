package usecase

import (
	"context"

	"github.com/pocket7878/spa_login_learning_backend/domain"
)

type TodoUsecaseImpl struct {
	repo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) *TodoUsecaseImpl {
	return &TodoUsecaseImpl{
		repo: repo,
	}
}

func (u *TodoUsecaseImpl) GetTodos(ctx context.Context, userID int64) ([]*domain.Todo, error) {
	todos, err := u.repo.GetTodos(ctx, userID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (u *TodoUsecaseImpl) Create(ctx context.Context, userID int64, description string) (*domain.Todo, error) {
	todo, err := u.repo.Create(ctx, userID, description)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *TodoUsecaseImpl) Delete(ctx context.Context, todoID int64) error {
	err := u.repo.Delete(ctx, todoID)
	return err
}
