package models

import "time"

type User struct {
	ID               int        `db:"id"`
	Username         string     `db:"username"`
	Email            string     `db:"email"`
	HashPassword     string     `db:"hash_password"`
	HashRefreshToken string     `db:"hash_refreshtoken"`
	Status           ItemStatus `db:"status"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
	DeletedAt        time.Time  `db:"deleted_at"`
}
