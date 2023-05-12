package models

import "time"

type Project struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Status      ItemStatus `db:"status"`
	OwnerID     int        `db:"owner_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   time.Time  `db:"deleted_at"`
}
