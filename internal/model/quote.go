package model

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author" binding:"required"`
	Quote  string `json:"quote" binding:"required"`
}
