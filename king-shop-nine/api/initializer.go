package api

import (
	"king-shop-nine/database"
	"king-shop-nine/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeApi() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/product/:id", getProductByID)
	router.POST("/product", postProduct)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	products := database.FetchProducts()

	c.IndentedJSON(http.StatusOK, products)

	// curl http://localhost:8080/products \
	//   --header "Content-Type: application/json" \
	//   --request "GET"
}

func getProductByID(c *gin.Context) {
	productID, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "product id provided is invalid"})
	}

	product := database.FetchProductByID(productID)
	c.IndentedJSON(http.StatusOK, product)

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})

	// curl http://localhost:8080/product/2
}

func postProduct(c *gin.Context) {
	var newProduct utils.Product

	// Call BindJSON to bind the received JSON to newProduct
	err := c.BindJSON(&newProduct)
	if err != nil {
		// log.Printf("ERROR => %v\n", err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "product data is invalid"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "product created"})
	// log.Printf("ALRIGHT\n")

	// Add the new product to the slice
	// utils.MockProducts = append(utils.MockProducts, newProduct)
	// c.IndentedJSON(http.StatusCreated, newProduct)

	// curl http://localhost:8080/product --include --header "Content-Type: application/json" --request "POST" --data '{"id": 35,"title": "Old school jeans","price": 34.99,"shortDescription": "Be cool, be oldschool!","description": "Old school style jeans that return you back to the good ol` days!","category": "unisex","subcategory": "all","image": "https://i.ibb.co/ZYW3VTp/brown-brim.png","dateAdded": "1682267143000","rating":{"rate": 4.6,"count": 15}}'
	// {"id": 35,"title": "Old school jeans","price": 34.99,"shortDescription": "Be cool, be oldschool!","description": "Old school style jeans that return you back to the good ol` days!","category": "unisex","subcategory": "all","image": "https://i.ibb.co/ZYW3VTp/brown-brim.png","dateAdded": "1682267143000","rating": {"rate": 4.6,"count": 15}}
}
