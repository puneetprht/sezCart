package controller

import (
	"strconv"
	"time"

	"goApi/database"

	"github.com/gin-gonic/gin"
)


type User struct {
	ID       	   uint64 	 `json:"id" gorm:"primary_key;auto_increment"`
	Name 				 string 	 `json:"name" gorm:"type:varchar(250)"`
	Username  	 string 	 `json:"username" gorm:"type:varchar(250)"`
	Password  	 string 	 `json:"password" gorm:"type:varchar(250)"`
	Cart_Id   	 uint64		 `json:"cart_id"`
	Token 			 string    `json:"token" gorm:"type:varchar(500)"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Cart struct {
	ID        	 uint64		 `json:"id" gorm:"primary_key;auto_increment"`
	User_Id      uint64 	 `json:"user_id"`
	Is_Purchased bool		   `json:"is_purchased"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Item struct {
	ID    	     uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name 	       string    `json:"name" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type CartItem struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Item_Id   	 uint64    `gorm:"primary_key;auto_increment" json:"item_id"`
	Cart_Id   	 uint64    `gorm:"primary_key;auto_increment" json:"cart_id"`
}

type Order struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Cart_Id   	 uint64		 `json:"cart_id"`
	User_Id   	 uint64		 `json:"user_id"`
	Created_At   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func UserPost(c *gin.Context) {
	var userLogin User
  if err := c.BindJSON(&userLogin); err != nil {
      return
  }

	userLogin.Created_At = time.Now() 
	result := database.DBConn.Create(&userLogin)
	c.JSON(200, result)
}

// token assignment left
func UserLogin(c *gin.Context) {
	var userLogin User
  if err := c.BindJSON(&userLogin); err != nil {
      return
  }
	userLogin.Created_At = time.Now() 
	result := database.DBConn.Create(&userLogin)
	c.JSON(200, result)
}

func UserList(c *gin.Context) {
	var users []User
	database.DBConn.Find(&users)
	c.JSON(200, users)
}

func ItemPost(c *gin.Context) {
	var newItem Item
  if err := c.BindJSON(&newItem); err != nil {
      return
  }
	newItem.Created_At = time.Now() 
	result := database.DBConn.Create(&newItem)
	c.JSON(200, result)
}

func ItemList(c *gin.Context) {
	var items []Item
	database.DBConn.Find(&items)
	c.JSON(200, items)
}

//cart post code
func CartPost(c *gin.Context) {
	var cartItem CartItem
  if err := c.BindJSON(&cartItem); err != nil {
      return
  }
	if (cartItem.Cart_Id == 0) {
		var newCart Cart
		var id = uint64(1)
		newCart.User_Id = id // replace with userId for loggedin user from token
		newCart.Created_At = time.Now()
		database.DBConn.Create(&newCart)
		cartItem.Cart_Id = newCart.ID
		var user User
		database.DBConn.First(&user, id)
		user.Cart_Id = newCart.ID
		database.DBConn.Save(&user)
	}
	if (cartItem.Cart_Id > 0 && cartItem.Item_Id > 0) {
		database.DBConn.Create(&cartItem)
	}
	c.JSON(200, cartItem)
}

//cart complete
func CartComplete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("cartId"))
	var cart Cart
	database.DBConn.First(&cart, id)
	cart.Is_Purchased = true
	database.DBConn.Save(&cart)

	var orderFinal Order
	orderFinal.Cart_Id = cart.ID
	orderFinal.User_Id = uint64(1)
	orderFinal.Created_At = time.Now()
	database.DBConn.Create(&orderFinal)

	c.JSON(200, orderFinal)
}

//I guess to be joined with order and users table to show cart and user info.
func CartList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("cartId"))
	var cartItems []CartItem
	if (id > 0) {
		database.DBConn.Where("cart_id = ?", id).Find(&cartItems)
	} else {
		database.DBConn.Find(&cartItems)	
	}	
	c.JSON(200, cartItems)
}

func OrderList(c *gin.Context) {
	var orders []Order
	database.DBConn.Find(&orders)
	c.JSON(200, orders)
}

