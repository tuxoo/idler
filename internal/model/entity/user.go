package entity

import "time"

type User struct {
	Id           int       `json:"-" db:"-"`
	Name         string    `json:"name" db:"name" binding:"required"`
	Email        string    `json:"email" db:"email" binding:"required"`
	Password     string    `json:"-" db:"-" binding:"required"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
	VisitedAt    time.Time `json:"lastVisitAt" db:"visited_at"`
}
