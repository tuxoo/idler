package dto

type DialogDTO struct {
	Name         string `json:"name"`
	FirstUserId  int    `json:"firstUserId" binding:"required"`
	SecondUserId int    `json:"secondUserId" binding:"required"`
}
