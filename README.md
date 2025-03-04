# What is this repo ?

### This repository demonstrates two go projects.
### The first project is the "server", a go web application which serves clients and allows them to Create, Read, Update, Delete and Get All products saved in an slice.

### The second project is the "client", a go web appication which has the same APIs as the "server" but the difference is that the "client" requests to the "server" to Create, Read, Update, Delete and Get All products. It acts as a proxy.

# How to run:
- open a shell
- cd server/
- go build -o main ./cmd/main.go
- ./main
- open another shell
- cd client/
- go build -o main ./cmd/main.go

# Server API Endpoints:
- GET http://localhost:3000/products (Getting All Products)
- GET http://localhost:3000/products/{id} (Getting a Single Product)
- POST http://localhost:3000/products (Creating a Product)
- PUT http://localhost:3000/products/{id} (Updating a Product)
- DELETE http://localhost:3000/products/{id} (Deleting a Product)

# Client API Endpoints:
- GET http://localhost:4000/products (Getting All Products)
- GET http://localhost:4000/products/{id} (Getting a Single Product)
- POST http://localhost:4000/products (Creating a Product)
- PUT http://localhost:4000/products/{id} (Updating a Product)
- DELETE http://localhost:4000/products/{id} (Deleting a Product)