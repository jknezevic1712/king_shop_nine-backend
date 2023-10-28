package utils

type User struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	AccountCreatedAt string `json:"accountCreatedAt"`
	Image            string `json:"image"`
}

type ProductRating struct {
	Rate  float32 `json:"rate" form:"rate" binding:"required"`
	Count int     `json:"count" form:"count" binding:"required"`
}
type Product struct {
	ID               int           `json:"id"`
	Title            string        `json:"title" form:"title" binding:"required"`
	Price            float32       `json:"price" form:"price" binding:"required"`
	ShortDescription string        `json:"shortDescription" form:"shortDescription" binding:"required"`
	Description      string        `json:"description" form:"description" binding:"required"`
	Category         string        `json:"category" form:"category" binding:"required"`
	Subcategory      string        `json:"subcategory" form:"subcategory" binding:"required"`
	Image            string        `json:"image" form:"image" binding:"required"`
	DateAdded        string        `json:"dateAdded" form:"dateAdded" binding:"required"`
	Rating           ProductRating `json:"rating" form:"rating" binding:"required"`
}
