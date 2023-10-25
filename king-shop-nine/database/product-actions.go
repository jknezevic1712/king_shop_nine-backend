package database

import (
	"encoding/json"
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

// Fetch all products
//
// @returns []Product
func FetchProducts() string {
	var products []utils.Product
	conn := ConnectToDB()

	rows, err := conn.Query(`SELECT * FROM "Products"`)
	if err != nil {
		log.Printf("FetchProducts: error executing query to the DB, %v\n", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product utils.Product
		if err = rows.Scan(&product.ID, &product.Title, &product.ShortDescription, &product.Description, &product.Category, &product.Subcategory, &product.Image, &product.DateAdded, &product.Rating.Rate, &product.Rating.Count); err != nil {
			log.Printf("FetchProducts: error while fetching products, %v\n", err)
		}

		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		log.Printf("FetchProducts: error while fetching products, %v\n", err)
	}

	marshalledRes, err := json.Marshal(products)
	if err != nil {
		log.Printf("FetchProducts: error while converting result to json, %v\n", err)
	}
	stringifiedJson := string(marshalledRes)

	conn.Close()
	return stringifiedJson
}

// Fetch product by ID
//
// @args productID of int64 type
//
// @returns Product
// func FetchProductByID(productID int64) utils.Product {
// 	var jsonData []byte
// for rows.Next() {
// 		if err := rows.Scan(&jsonData); err != nil {
// 			log.Printf("FetchProducts: error while fetching products, %v\n", err)
// 		}
// 	}

// 	var product utils.Product
// 	conn := ConnectToDB()

// 	row := conn.QueryRow(`SELECT * FROM "Products" WHERE id = $1`, productID)

// 	if err := row.Scan(&product.ID, &product.Title, &product.ShortDescription, &product.Description, &product.Category, &product.Subcategory, &product.Image, &product.DateAdded, &product.Rating.Rate, &product.Rating.Count); err != nil {
// 		log.Printf("FetchProductByID: error while fetching product with id %c, %v\n", productID, err)
// 	}

// 	log.Printf("FetchProductByID: fetched product => %v\n", product)
// 	conn.Close()
// 	return product
// }
