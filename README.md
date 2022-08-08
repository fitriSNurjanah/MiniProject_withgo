# Golang Restful API using GORM ORM (Postgres), Gin, JWT

## this is a mini prokject with go

# **FOLDER STRUCTURE**

miniproject  
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
│ ├─ loginRequest.go  
│ ├─ loginResponse.go  
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
│ ├─ productService.go  
│ ├─ productService_test.go  
│ └─ userService.go  
├─ go.mod  
├─ go.sum  
├─ main.go  
└─ README.md

## Package Used to create REST API

go get -u github.com/gin-gonic/gin
go get -u github.com/joho/godotenv
go get -u gorm.io/gorm
gorm.io/driver/postgres
go get -u go.uber.org/zap
go get -u github.com/golang-jwt/jwt/v4

## Setting Configure File

SERVER_ADDRESS=localhost # Server Addreaa
SERVER_PORT=8080 # Server Port
DB_USER=postgres # Postgres
DB_PASSWD=fitri123 # Database Password
DB_ADDR=localhost # Database Host
DB_PORT=5432 # Database Port
DB_NAME=projects # Database Name

# **Running Project**

Go run .

# Create Table Database

CREATE TABLE products
( id serial NOT NULL,
merk char(50) NOT NULL,
price int,
description char(50)
);

INSERT INTO products
(id, merk, price, description)
VALUES
(1,'samsung',500,'samsung z flip');

CREATE TABLE users
( id int NOT NULL,
username char(50) NOT NULL,
password char(50),
CONSTRAINT users_pk PRIMARY KEY (id)
);

INSERT INTO users
(id, username, password)
VALUES
(1,'admin','admin');

# API Endpoint
    router.GET("/products", ch.getAllProduct)  search with Get All Product
    router.GET("/products/:id", ch.getProductID) search with id
    router.POST("/products", ch.createProduct) create new product
    router.DELETE("/products/:id", ch.DeleteProduct) delete product
    router.PUT("/products/:id", ch.UpdateProduct) update product
    router.POST("/users", ah.registerUser) create new user

## Create New Product
1. localhost:8000/products
{
    "merk" : "Oppo",
    "price" : 15,
    "description" : "Oppo A21"

}

output
{
    "id": 7,
    "merk": "Oppo",
    "price": 15,
    "description": "Oppo A21"
}
