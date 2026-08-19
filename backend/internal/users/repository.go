package users

import (
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := "SELECT (id, username, email) FROM users WHERE email = $1"
	var user User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	query := "SELECT (id, username, email) FROM users WHERE username = $1"
	var user User
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *Repository) Create(ctx context.Context, username, email, hashedPassword string) (*User, error) {

	var user User

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email, created_at"
	err := r.db.QueryRow(query, username, email, hashedPassword).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
