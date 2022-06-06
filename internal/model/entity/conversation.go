package entity

type Conversation struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Owner int    `json:"owner" binding:"required"`
	//Participants []dto.UserDTO `json:"participants" binding:"required"`
}
