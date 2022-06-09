package entity

import "time"
import . "github.com/google/uuid"

type User struct {
	Id           UUID      `json:"-" db:"id"`
	Name         string    `json:"name" db:"name" binding:"required"`
	Email        string    `json:"email" db:"email" binding:"required"`
	Password     string    `json:"-" db:"-" binding:"required"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
	VisitedAt    time.Time `json:"lastVisitAt" db:"visited_at"`
	IsConfirmed  bool      `json:"-" db:"is_confirmed"`
}
