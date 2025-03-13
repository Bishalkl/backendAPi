package user

import (
	"database/sql"
	"fmt"

	"github.com/bishalkl/backendAPi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// Use QueryRow since we expect a single record
	row := s.db.QueryRow("SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?", email)

	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
