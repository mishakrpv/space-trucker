package user

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID `db:"user_id"`
	Email   string    `db:"email"`
	PwdHash string    `db:"password_hash"`
}
