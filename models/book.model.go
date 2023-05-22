package models

type Book struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Rating      float64 `json:"rating"`
	Price       uint    `json:"price"`
	ImageUrl    string  `json:"image_url"`
}
