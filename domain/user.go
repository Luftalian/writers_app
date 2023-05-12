package domain

import (
	"time"
)

type User struct {
	UserID    UUID      `json:"user_id" db:"user_id" form:"user_id"`
	Name      string    `json:"name" db:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at" form:"created_at"`
}

type Users []User

type UserCreate struct {
	UserID UUID `json:"user_id" db:"user_id"`
	TextID UUID `json:"text_id" db:"text_id"`
}

type UserCreates []UserCreate

type UserLike struct {
	UserID UUID `json:"user_id" db:"user_id"`
	TextID UUID `json:"text_id" db:"text_id"`
}

type UserLikes []UserLike
