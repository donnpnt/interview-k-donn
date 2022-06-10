package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zipzoft/interview-k-donn/service"
)

func main() {

	// Initial gin framework
	app := gin.Default()
	app.LoadHTMLGlob("views/*")

	// Initial service
	productService := service.NewProductService()

	app.GET("/", func(c *gin.Context) {
		products, err := productService.All()

		// Show products page
		c.HTML(http.StatusOK, "products.tmpl", gin.H{
			"products": products,
			"err":      err,
		})
	})

	// Run server
	app.Run(":8080")
}
