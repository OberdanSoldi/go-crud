# Simple CRUD API
A lightweight RESTful API built with Go, using Gin framework and SQLite database.
## Overview
This project implements a simple CRUD (Create, Read, Update, Delete) API for managing books. It follows a clean architecture pattern with clear separation of concerns between adapters, models, and use cases.
## Technologies Used
- Go 1.24
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [SQLite](https://github.com/glebarez/sqlite) - Database
## Project Structure
``` 
simple-crud/
├── src/
│   ├── adapter/
│   │   ├── controller/    - HTTP request handlers
│   │   ├── db/            - Database connection and initialization
│   │   └── repository/    - Data access layer
│   ├── model/
│   │   ├── entity/        - Entities
│   │   ├── request/       - Request models
│   │   └── response/      - Response models
│   ├── usecase/
│   │   ├── converter/     - Data conversion between layers
│   │   ├── interfaces/    - Implementation contracts
│   │   └── service/       - Business logic
│   └── server.go          - Server initialization
└── main.go                - Application entry point
```
## API Routes
The application exposes the following RESTful endpoints for managing books:

| Method | Path | Description |
| --- | --- | --- |
| GET | `/books` | Retrieve all books |
| GET | `/books/:id` | Retrieve a book by ID |
| POST | `/books` | Create a new book |
| PUT | `/books/:id` | Update an existing book |
| DELETE | `/books/:id` | Delete a book by ID |
## Getting Started
### Prerequisites
- Go 1.24 or later

### Installation
1. Clone the repository
``` bash
git clone https://github.com/yourusername/simple-crud.git
cd simple-crud
```
1. Install dependencies
``` bash
go mod download
```
1. Run the application
``` bash
go run main.go
```
The server will start on port 8080.
## Usage Example
### Create a book
``` bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "The Go Programming Language", "author": "Alan A. A. Donovan", "year": 2015}'
```
### Get all books
``` bash
curl http://localhost:8080/books
```
### Get a specific book
``` bash
curl http://localhost:8080/books/1
```
### Update a book
``` bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "The Go Programming Language", "author": "Alan A. A. Donovan & Brian W. Kernighan", "published_date": "2006-01-02T15:04:05Z"}'
```
### Delete a book
``` bash
curl -X DELETE http://localhost:8080/books/1
```
