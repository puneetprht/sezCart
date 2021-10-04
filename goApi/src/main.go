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
	
	fmt.Printf("SezCart API.")
	routes.POST("/user/create", controller.UserPost)
	routes.POST("/user/login", controller.UserLogin)
	routes.GET("/user/list", controller.UserList)
	routes.POST("/item/create", controller.ItemPost)
	routes.GET("/item/list", controller.ItemList)
	routes.POST("/cart/add", controller.CartPost)
	routes.PUT("/cart/:cartId/complete", controller.CartComplete)
	routes.GET("/cart/list", controller.CartList)
	routes.GET("/order/list", controller.OrderList)

	routes.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:root@/ecom?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&controller.User{}, &controller.Cart{}, &controller.Item{}, &controller.CartItems{}, &controller.Order{})
}
