# Golang Restful API using GORM ORM (Postgres), Gin, JWT

### this is a mini prokject with go

## **FOLDER STRUCTURE**

miniproject  
├─ API DOC  
│ └─ MiniProject.postman_collection.json  
├─ app  
│ ├─ app.go  
│ ├─ productHandler.go  
│ ├─ productHandler_test.go  
│ └─ userHandler.go  
├─ domain  
│ ├─ product.go  
│ ├─ productRepositoryDB.go  
│ ├─ productRepositoryDB_test.go  
│ ├─ user.go  
│ └─ userRepositoryDB.go  
├─ dto  
│ ├─ productPagination.go  
│ ├─ productRequest.go  
│ ├─ productResponse.go  
│ └─ userRequest.go  
├─ errs  
│ └─ errors.go  
├─ helper  
│ └─ paginations.go  
├─ logger  
│ └─ logger.go  
├─ service  
│ ├─ authService.go  
│ ├─ productService.go  
│ └─ userService.go  
├─ setupDB  
│ └─ setupDB.go  
├─ go.mod  
├─ go.sum  
├─ main.go  
└─ README.md

## Package Used to create REST API

go get -u github.com/gin-gonic/gin<br>
go get -u github.com/joho/godotenv<br>
go get -u gorm.io/gorm<br>
gorm.io/driver/postgres<br>
go get -u go.uber.org/zap<br>
go get -u github.com/golang-jwt/jwt/v4<br><br>

## Setting Configure File

SERVER_ADDRESS=localhost # Server Address<br>
SERVER_PORT=8080 # Server Port<br>
DB_USER=postgres # Postgres<br>
DB_PASSWD=fitri123 # Database Password<br>
DB_ADDR=localhost # Database Host<br>
DB_PORT=5432 # Database Port<br>
DB_NAME=projects # Database Name<br>

## Running Project

**Go run .**

## Create Table Database

CREATE TABLE products<br>
( id serial NOT NULL,<br>
merk char(50) NOT NULL,<br>
price int,<br>
description char(50)<br>
);<br><br>

INSERT INTO products<br>
(id, merk, price, description)<br>
VALUES<br>
(1,'samsung',500,'samsung z flip');<br><br><br>

CREATE TABLE users<br>
( id int NOT NULL,<br>
username char(50) NOT NULL,<br>
password char(50),<br>
CONSTRAINT users_pk PRIMARY KEY (id)<br>
);<br><br>

INSERT INTO users<br>
(id, username, password)<br>
VALUES<br>
(1,'admin','admin');<br>

## API Endpoint

    router.GET("/products", ch.getAllProduct)  search with Get All Product
    router.GET("/products/:id", ch.getProductID) search with id
    router.POST("/products", ch.createProduct) create new product
    router.DELETE("/products/:id", ch.DeleteProduct) delete product
    router.PUT("/products/:id", ch.UpdateProduct) update product
    router.POST("/users", ah.registerUser) create new user

## Create New Product

localhost:8000/products<br>
{<br>
"merk" : "Oppo",<br>
"price" : 15,<br>
"description" : "Oppo A21"<br>
}<br>

**output**<br>
{<br>
"id": 7,<br>
"merk": "Oppo",<br>
"price": 15,<br>
"description": "Oppo A21"<br>
}
