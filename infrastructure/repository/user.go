package repository

import (
	"context"
	"database/sql"

	"github.com/pocket7878/spa_login_learning_backend/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (u *UserRepositoryImpl) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT id,email FROM users`

	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	result := new(domain.User)
	err = rows.Scan(&result.ID, &result.Email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserRepositoryImpl) Store(ctx context.Context, user *domain.User) error {
	query := `INSERT users SET email=?`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, user.Email)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = lastID

	return nil
}
