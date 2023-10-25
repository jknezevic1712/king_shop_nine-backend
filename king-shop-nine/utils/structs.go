package utils

type User struct {
	ID               string
	Name             string
	Email            string
	AccountCreatedAt string
	Image            string
}

type ProductRating struct {
	Rate  float32
	Count int
}
type Product struct {
	ID               int
	Title            string
	Price            float32
	ShortDescription string
	Description      string
	Category         string
	Subcategory      string
	Image            string
	DateAdded        string
	Rating           ProductRating
}
