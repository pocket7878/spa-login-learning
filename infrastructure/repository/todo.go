package repository

import (
	"context"
	"database/sql"
	"os"

	"github.com/lib/pq"
	"github.com/pocket7878/spa_login_learning_backend/domain"
)

type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepository() (*TodoRepositoryImpl, error) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	return &TodoRepositoryImpl{db}, nil
}

func (r *TodoRepositoryImpl) GetTodos(ctx context.Context, userID int64) ([]*domain.Todo, error) {
	query := `SELECT id, user_id, description FROM todos WHERE user_id = $1`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*domain.Todo, 0)
	for rows.Next() {
		todo := new(domain.Todo)
		err := rows.Scan(&todo.ID, &todo.UserID, &todo.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, todo)
	}

	return result, nil
}

func (r *TodoRepositoryImpl) GetTodo(ctx context.Context, todoID int64) (*domain.Todo, error) {
	query := `SELECT id, user_id, description FROM todos WHERE id = $1`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, todoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := new(domain.Todo)
	if rows.Next() {
		err := rows.Scan(&result.ID, &result.UserID, &result.Description)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return nil, nil
	}
}

func (r *TodoRepositoryImpl) Create(ctx context.Context, userID int64, description string) (*domain.Todo, error) {
	query := `INSERT INTO todos (user_id, description) VALUES ($1, $2) RETURNING id`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var result = &domain.Todo{
		UserID:      userID,
		Description: description,
	}
	err = stmt.QueryRowContext(ctx, userID, description).Scan(&result.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TodoRepositoryImpl) Delete(ctx context.Context, todoID int64) error {
	query := `DELETE FROM todos WHERE id = $1`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, todoID)
	if err != nil {
		return err
	}

	return nil
}
