package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"goApi/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authCustomClaims struct {
	Name 				string 		 `json:"name"`
	jwt.StandardClaims
}

type AuthToken struct {
	Token 			string 		 `json:"token"`
}

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
	result := database.DBConn.First(&userLogin, "username = ?", userLogin.Username)
	if(result.RowsAffected == 1){
		c.JSON(500, gin.H{
			"message": "User already exists",
		})	
		return
	}
	userLogin.Created_At = time.Now() 
	result = database.DBConn.Create(&userLogin)
	c.JSON(200, result)
}

// token assignment left
func UserLogin(c *gin.Context) {
	var userLogin User
  if err := c.BindJSON(&userLogin); err != nil {
      return
  }
	result := database.DBConn.First(&userLogin, "username = ? AND password >= ?", userLogin.Username, userLogin.Password)
	if(result.RowsAffected == 0){
		c.JSON(500, gin.H{
			"message": "username/password is invalid!",
		})	
		return 
	}
	token := generateToken(userLogin.Username)
	userLogin.Token = token
	database.DBConn.Save(&userLogin)
	c.JSON(200, userLogin)
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
	auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusForbidden, "No Authorization header provided")
			return
		}
	tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == auth {
			c.JSON(403, "Could not find bearer token in Authorization header")
			return
		}
	tokenOutput, err := validateToken(tokenString)
	if err != nil {
		c.JSON(403, "Invalid authentication token")
	}
	if tokenOutput.Valid == false {
		c.JSON(403, "Invalid authentication token")
	}
	var user User
	database.DBConn.First(&user, " token = ? ", tokenString)
	
	var cartItem CartItem
  if err := c.BindJSON(&cartItem); err != nil {
      return
  }
	if (cartItem.Cart_Id == 0) {
		var newCart Cart
		newCart.User_Id = user.ID 
		newCart.Created_At = time.Now()
		database.DBConn.Create(&newCart)
		cartItem.Cart_Id = newCart.ID
		var user User
		database.DBConn.First(&user, user.ID)
		user.Cart_Id = newCart.ID
		database.DBConn.Save(&user)
	}
	if (cartItem.Cart_Id > 0 && cartItem.Item_Id > 0) {
		database.DBConn.Create(&cartItem)
	}
	c.JSON(200, cartItem)
}

func CartComplete(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusForbidden, "No Authorization header provided")
			return
		}
	tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == auth {
			c.JSON(403, "Could not find bearer token in Authorization header")
			return
		}
	tokenOutput, err := validateToken(tokenString)
	if err != nil {
		c.JSON(403, "Invalid authentication token")
	}
	if tokenOutput.Valid == false {
		c.JSON(403, "Invalid authentication token")
	}
	var user User
	database.DBConn.First(&user, " token = ? ", tokenString)

	id, _ := strconv.Atoi(c.Param("cartId"))
	var cart Cart
	database.DBConn.First(&cart, id)
	cart.Is_Purchased = true
	database.DBConn.Save(&cart)

	var orderFinal Order
	orderFinal.Cart_Id = cart.ID
	orderFinal.User_Id = user.ID
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

func ValidateUserToken(c *gin.Context) {
	var userToken AuthToken
  if err := c.BindJSON(&userToken); err != nil {
      return
  }
	tokenOutput, err := validateToken(userToken.Token)
	if err != nil {
		c.JSON(403, "Invalid authentication token")
	}
	if tokenOutput.Valid == false {
		c.JSON(403, "Invalid authentication token")
	}
	var user User
	database.DBConn.First(&user, " token = ? ", userToken.Token)
	c.JSON(200, user)
}

func generateToken(email string) string {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "sezcart",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte("3663083757"))
	if err != nil {
		panic(err)
	}
	return t
}

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte("3663083757"), nil
	})
}