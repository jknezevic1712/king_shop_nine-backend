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
	products, err := database.FetchProducts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "products could not be retrieved", "details": err})
	}

	c.JSON(http.StatusOK, products)

	// curl http://localhost:8080/products \
	//   --header "Content-Type: application/json" \
	//   --request "GET"
}

func getProductByID(c *gin.Context) {
	productID, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "product id provided is invalid", "details": err})
	}

	product, err := database.FetchProductByID(productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "product with the specified ID could not be retrieved", "details": err})
	}

	c.JSON(http.StatusOK, product)

	// curl http://localhost:8080/product/2
}

func postProduct(c *gin.Context) {
	newProduct := utils.Product{}

	// Call BindJSON to bind the received JSON to newProduct
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "product data is invalid", "details": err})
		return
	}

	if err := database.AddProduct(newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "product unsuccessfully created", "details": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product successfully created", "details": nil})

	// curl http://localhost:8080/product --include --header "Content-Type: application/json" --request "POST" --data '{"id": 35,"title": "Old school jeans","price": 34.99,"shortDescription": "Be cool, be oldschool!","description": "Old school style jeans that return you back to the good ol` days!","category": "unisex","subcategory": "all","image": "https://i.ibb.co/ZYW3VTp/brown-brim.png","dateAdded": "1682267143000","rating":{"rate": 4.6,"count": 15}}'
	// {"id": 35,"title": "Old school jeans","price": 34.99,"shortDescription": "Be cool, be oldschool!","description": "Old school style jeans that return you back to the good ol` days!","category": "unisex","subcategory": "all","image": "https://i.ibb.co/ZYW3VTp/brown-brim.png","dateAdded": "1682267143000","rating": {"rate": 4.6,"count": 15}}
}
