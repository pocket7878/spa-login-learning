package domain

import "context"

type Todo struct {
	ID          int64
	UserID      int64
	Description string
}

type TodoUsecase interface {
	GetTodos(ctx context.Context, userID int64) ([]*Todo, error)
	Create(ctx context.Context, userID int64, description string) (*Todo, error)
	Delete(ctx context.Context, todoID int64) error
}

type TodoRepository interface {
	GetTodos(ctx context.Context, userID int64) ([]*Todo, error)
	Create(ctx context.Context, userID int64, description string) (*Todo, error)
	Delete(ctx context.Context, todoID int64) error
}
