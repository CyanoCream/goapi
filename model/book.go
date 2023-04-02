package model

import (
	"time"
)

type Books struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"name_book" db:"title"`
	Author      string `json:"author" db:"author"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"created_at"`
}
