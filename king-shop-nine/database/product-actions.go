package database

import (
	"king-shop-nine/utils"
	"log"
)

// Add product to the DB
//
// @args newProduct of Product struct type
//
// @returns error
func AddProduct(newProduct utils.Product) error {
	conn := ConnectToDB()

	q := `
		INSERT INTO "Products" (title, "shortDescription", description, category, subcategory, image, "dateAdded", "ratingRate", "ratingCount")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := conn.Exec(q, &newProduct.Title, &newProduct.ShortDescription, &newProduct.Description, &newProduct.Category, &newProduct.Subcategory, &newProduct.Image, &newProduct.DateAdded, &newProduct.Rating.Rate, &newProduct.Rating.Count)
	if err != nil {
		log.Printf("AddProduct: error while inserting new product, %v\n", err)
		return err
	}

	log.Println("AddProduct: succesfully inserted new product")

	conn.Close()
	return nil
}
