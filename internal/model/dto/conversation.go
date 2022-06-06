package dto

type ConversationDTO struct {
	Name  string `json:"name" binding:"required"`
	Owner int    `json:"owner"`
	//Participant UserDTO `json:"participants" binding:"required"`
}
