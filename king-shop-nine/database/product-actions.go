package database

import (
	"king-shop-nine/utils"
	"log"
)

// Add product to the DB
//
// @args newProduct Product
//
// @returns error
func AddProduct(newProduct utils.Product) error {
	conn := ConnectToDB()

	q := `
		INSERT INTO "Products" (title, "shortDescription", description, category, subcategory, image, "dateAdded", "ratingRate", "ratingCount")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := conn.Exec(q, newProduct.Title, newProduct.ShortDescription, newProduct.Description, newProduct.Category, newProduct.Subcategory, newProduct.Image, newProduct.DateAdded, newProduct.Rating.Rate, newProduct.Rating.Count)
	if err != nil {
		log.Printf("AddProduct: error while inserting new product, %v\n", err)
		return err
	}

	log.Println("AddProduct: succesfully inserted new product")

	conn.Close()
	return nil
}

// Fetch all products
//
// @returns []utils.Product
func FetchProducts() ([]utils.Product, error) {
	var products []utils.Product
	conn := ConnectToDB()

	rows, err := conn.Query(`SELECT * FROM "Products"`)
	if err != nil {
		log.Printf("FetchProducts: error executing query to the DB, %v\n", err)
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product utils.Product
		if err = rows.Scan(&product.ID, &product.Title, &product.ShortDescription, &product.Description, &product.Category, &product.Subcategory, &product.Image, &product.DateAdded, &product.Rating.Rate, &product.Rating.Count); err != nil {
			log.Printf("FetchProducts: error while fetching products, %v\n", err)
			return products, err
		}

		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		log.Printf("FetchProducts: error while fetching products, %v\n", err)
		return products, err
	}

	conn.Close()
	return products, nil
}

// Fetch product by ID
//
// @args productID int
//
// @returns (utils.Product, error)
func FetchProductByID(productID int) (utils.Product, error) {
	var product utils.Product
	conn := ConnectToDB()

	row := conn.QueryRow(`SELECT * FROM "Products" WHERE id = $1`, productID)

	if err := row.Scan(&product.ID, &product.Title, &product.ShortDescription, &product.Description, &product.Category, &product.Subcategory, &product.Image, &product.DateAdded, &product.Rating.Rate, &product.Rating.Count); err != nil {
		log.Printf("FetchProductByID: error while fetching product with id %c, %v\n", productID, err)
		return product, err
	}

	conn.Close()
	return product, nil
}
