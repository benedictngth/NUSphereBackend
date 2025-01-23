package users

import "time"

type User struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	PublicID     string    `db:"public_id"`
}

type UserPublic struct {
	ID       string `db:"public_id"`
	Username string `db:"username"`
}
