package main

import (
	"fmt"

	"goApi/controller"
	"goApi/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {	
	initDatabase()
	defer database.DBConn.Close()

	routes := gin.Default()

	routes.Use(CORS())
	
	fmt.Printf("SezCart API.")
	routes.POST("/user/create", controller.UserPost)
	routes.POST("/user/login", controller.UserLogin)
	routes.GET("/user/list", controller.UserList)
	routes.POST("/token/validate", controller.ValidateUserToken)
	routes.POST("/item/create", controller.ItemPost)
	routes.GET("/item/list", controller.ItemList)
	routes.POST("/cart/add", controller.CartPost)
	routes.POST("/cart/:cartId/complete", controller.CartComplete)
	routes.GET("/cart/list", controller.CartList)
	routes.GET("/order/list", controller.OrderList)

	routes.Run(":4000") // listen on localhost:4000 for dev
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:root@/ecom?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&controller.User{}, &controller.Cart{}, &controller.Item{}, &controller.Cartitem{}, &controller.Order{})
}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
			c.Writer.Header().Set("Content-Type", "application/json")
			if c.Request.Method != "OPTIONS" {
				c.Next()
			} else {
				c.AbortWithStatus(200)
			}
	}
}