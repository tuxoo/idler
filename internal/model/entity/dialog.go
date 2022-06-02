package entity

import "time"

type Dialog struct {
	Id           int       `json:"-" db:"id"`
	Name         string    `json:"name" db:"name" binding:"required"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	LastMessage  time.Time `json:"lastMessage" db:"last_message"`
	FirstUserId  int       `json:"first_user_id" db:"first_user_id"`
	SecondUserId int       `json:"second_user_id" db:"second_user_id"`
}
