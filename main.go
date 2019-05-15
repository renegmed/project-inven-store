package main 

import (
	"net/http"
	"os"
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context){
		c.String(http.StatusOK, "pong")
	}) 

	engine.GET("/api/products", func(c *gin.Context){
		c.JSON(http.StatusOK, AllProducts())
	})

	engine.GET("/api/reset", func(c *gin.Context){
		ResetQty()
		c.Status(http.StatusOK) 
	})

	engine.PUT("/api/reduce/:sku/:qty", func(c *gin.Context){
		sku := c.Params.ByName("sku")
		qty := c.Params.ByName("qty")
		product, found := GetProduct(sku)
		fmt.Println("Point 0")
		if found {
			qtyToReduce, err := strconv.Atoi(qty)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound) 
			} else {
				fmt.Println("Point 1")
				product.Quantity = product.Quantity - qtyToReduce
				updated := UpdateProduct(sku, product)
				if updated { 
					fmt.Printf("Point 2 \n %v", product)
					c.Status(http.StatusOK)
				} else {
					fmt.Println("Point 3")
					c.Status(http.StatusNotFound)
				}
			}	  
		} else {
			fmt.Println("Point 9")
			c.AbortWithStatus(http.StatusNotFound)
		} 
	})


	engine.PUT("/api/increase/:sku/:qty", func(c *gin.Context){
		sku := c.Params.ByName("sku")
		qty := c.Params.ByName("qty")
		product, found := GetProduct(sku)
		fmt.Println("Point 0")
		if found {
			qtyToAdd, err := strconv.Atoi(qty)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound) 
			} else {
				fmt.Println("Point 1")
				product.Quantity = product.Quantity + qtyToAdd
				updated := UpdateProduct(sku, product)
				if updated { 
					fmt.Printf("Point 2 \n %v", product)
					c.Status(http.StatusOK)
				} else {
					fmt.Println("Point 3")
					c.Status(http.StatusNotFound)
				}
			}	  
		} else {
			fmt.Println("Point 9")
			c.AbortWithStatus(http.StatusNotFound)
		} 
	})

	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}