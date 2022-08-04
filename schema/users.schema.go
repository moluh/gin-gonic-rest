package schema

import "time"

type SchemaUser struct {
	ID        string    `json:"id" validate:"uuid"`
	Name      string    `json:"name" validate:"required,lowercase"`
	Surname   string    `json:"surname" validate:"required,lowercase"`
	Email     string    `json:"email" validate:"required,email"`
	Token     string    `json:"token" validate:"required"`
	Password  string    `json:"password" validate:"required,gte=8"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
