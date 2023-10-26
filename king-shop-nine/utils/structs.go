package utils

type User struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	AccountCreatedAt string `json:"accountCreatedAt"`
	Image            string `json:"image"`
}

type ProductRating struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}
type Product struct {
	ID               int           `json:"id"`
	Title            string        `json:"title"`
	Price            float32       `json:"price"`
	ShortDescription string        `json:"shortDescription"`
	Description      string        `json:"description"`
	Category         string        `json:"category"`
	Subcategory      string        `json:"subcategory"`
	Image            string        `json:"image"`
	DateAdded        string        `json:"dateAdded"`
	Rating           ProductRating `json:"rating"`
}
