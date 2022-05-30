package model

import "time"

type User struct {
	Id           int       `json:"-" db:"id"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Username     string    `json:"username" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	RegisteredAt time.Time `json:"registeredAt" bson:"registeredAt"`
	LastVisitAt  time.Time `json:"lastVisitAt" bson:"lastVisitAt"`
}
