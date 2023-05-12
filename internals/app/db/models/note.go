package models

import "time"

type Note struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Content     string     `db:"content"`
	Format      NoteFormat `db:"format"`
	Status      ItemStatus `db:"status"`
	UserID      int        `db:"user_id"`
	ProjectID   int        `db:"project_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   time.Time  `db:"deleted_at"`
}
