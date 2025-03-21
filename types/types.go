package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName`
	LastName  string    `json:"lastName`
	Email     string    `json: "email`
	Password  string    `json: "-"`
	CreatedAt time.Time `json: "createdAt"`
}

// payload types of register
type RegisterUserPayload struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName`
	Email     string `json: "email`
	Password  string `json: "password"`
}
