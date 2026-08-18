package user

import "database/sql"

type Repository struct {
	db *sql.DB
}

func (r *Repository) Insert(username, email, hashedPassword string) (User, error) {

	var u User

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email, created_at"
	err := r.db.QueryRow(query, username, email, hashedPassword).Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return u, nil

}
