package entity

import "time"

type User struct {
	Id           int       `json:"-" db:"id"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	RegisteredAt time.Time `json:"registeredAt" bson:"registeredAt"`
	VisitedAt    time.Time `json:"lastVisitAt" bson:"lastVisitAt"`
}
