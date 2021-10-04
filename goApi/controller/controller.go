package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


type User struct {
	ID       	   uint64 	 `json:"id" gorm:"primary_key;auto_increment"`
	Name 				 string 	 `json:"name" gorm:"type:varchar(250)"`
	Username  	 string 	 `json:"username" gorm:"type:varchar(250)"`
	Password  	 string 	 `json:"password" gorm:"type:varchar(250)"`
	Cart   			 Cart 		 `gorm:"embedded;embeddedPrefix:cart_"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Cart struct {
	ID        	 uint64		 `json:"id" gorm:"primary_key;auto_increment"`
	User_Id      string 	 `json:"user_id"`
	Is_Purchased string		 `json:"is_purchased"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Item struct {
	ID    	     uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name 	       string    `json:"name" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type CartItems struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ItemId   		 uint64    `gorm:"primary_key;auto_increment" json:"itemId"`
	CartId   		 uint64    `gorm:"primary_key;auto_increment" json:"cartId"`
}

type Order struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Cart   			 Cart      `gorm:"embedded;embeddedPrefix:cart_"`
	User   			 User 		 `gorm:"embedded;embeddedPrefix:user_"`
	Is_Purchased string    `json:"is_purchased"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func UserPost(c *gin.Context) {
	var userLogin User
  if err := c.BindJSON(&userLogin); err != nil {
      return
  }
	fmt.Println("user posting function.")
	
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserLogin(c *gin.Context) {
	var userLogin User
  if err := c.BindJSON(&userLogin); err != nil {
      return
  }
	userLogin.Created_At = time.Now() 
	c.JSON(200, userLogin)
}

func UserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"nbame": "pong",
	})
}

func ItemPost(c *gin.Context) {
	var newItem Item
  if err := c.BindJSON(&newItem); err != nil {
      return
  }
	newItem.Created_At = time.Now() 
	c.JSON(200, newItem)
}

func ItemList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CartPost(c *gin.Context) {
	var cartItem CartItems
  if err := c.BindJSON(&cartItem); err != nil {
      return
  }
	c.JSON(200, cartItem)
}

func CartComplete(c *gin.Context) {
	name := c.Param("cartId")
	c.JSON(200, gin.H{
		"name": name,
	})
}

func CartList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func OrderList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

