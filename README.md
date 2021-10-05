# sezCart
 ECommerce Website Project

![Go](https://github.com/waznico/go-book-api/workflows/Go/badge.svg)

# Go React based Demo Project
Here a simple REST API has been created inorder to process transactions from the frontend. The REST API service is written in Go with Gin Framework and GORM ORM framework. The Frontend is based upon Next.js which is a lightweight react Framework. Code has been deployed on a private AWS EC2 server, Database MYSQL has been used.

# Usage
The API application opens a socket at the port 4000.

The following endpoints are currently exposed:

| Endpoint                     | Description                         |
|------------------------------|-------------------------------------|
| /user/create   		(POST) | Creates new user, email Id unique.  |
| /user/login	        (POST) | Userlogin, return user & token.     |
| /user/list            (GET)  | Fetch of all the users              |
| /token/validate       (POST) | To validate token on login          |
| /item/create          (POST) | Creates a new Item                  |
| /item/list   			(GET)  | Fetch all the Items                 |
| /cart/add   			(POST)*| Adds item to cart(Auth based)       |
| /cart/:cartId/complete(POST)*| Converts cart to Order              |
| /cart/list?cartId=id	(GET)  | Fetches items in a Cart(opt. CartId)|
| /order/list?userId=id (GET)  | Fetches all orders,(optional userId)|

Thanks
