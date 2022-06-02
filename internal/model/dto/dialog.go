package dto

type DialogDTO struct {
	Name     string `json:"name" binding:"required"`
	FirstId  int    `json:"firstId" binding:"required"`
	SecondId int    `json:"secondId" binding:"required"`
}
