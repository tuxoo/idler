package entity

import "time"

type Dialog struct {
	Id           int       `json:"-" db:"id"`
	Name         string    `json:"name" binding:"required"`
	CreatedAt    time.Time `json:"createdAt"`
	LastMessage  time.Time `json:"lastMessage"`
	FirstUserId  int       `json:"first_user_id" db:"first_user_is"`
	SecondUserId int       `json:"second_user_id" db:"second_user_is"`
}
