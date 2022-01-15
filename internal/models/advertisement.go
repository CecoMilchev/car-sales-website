package models

type Advertisement struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Power       uint   `json:"power"`
	Price       uint   `json:"price"`
	UserID      int    `json:"userID"`
	CarID       int    `json:"carID"`
}
