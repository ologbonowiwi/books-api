# bookstore-api
Books management API with Gin and Gorm

## Table of contents
- [bookstore-api](#bookstore-api)
  - [Table of contents](#table-of-contents)
  - [Description](#description)
  - [Installation](#installation)
  - [Requirements](#requirements)
    - [For hot reload:](#for-hot-reload)
  - [Running the app](#running-the-app)
    - [Running with hot reload](#running-with-hot-reload)
  - [Swagger](#swagger)
  - [Author](#author)
  - [License](#license)

## Installation
To install the app you need to run the following command:

```bash
git clone https://github.com/ologbonowiwi/bookstore-api.git
cd bookstore-api
go mod download
```

## Requirements
- Go 1.19

### For hot reload:
- [air](https://github.com/cosmtrek/air)

## Running the app
To run the app you need to run the following command:

```bash
go run ./cmd/main/main.go
```
or
```bash
go build -o bookstore-api ./cmd/main/main.go
./bookstore-api
```

### Running with hot reload
To run the app with hot reload you need to run the following command:

```bash
air
```

## Swagger
After running the app, you can access the swagger documentation at the following link: http://localhost:8080/swagger/index.html

To generate the swagger documentation you need to run the following command:

```bash
swag init -d ./cmd/main,./pkg
```

## Author
My name is [Wesley Matos](https://github.com/ologbonowiwi), and you can find me here (at GitHub 😁), on my [LinkedIn](https://www.linkedin.com/in/ologbonowiwi/) or by [e-mail](mailto:ologbonowiwi520@gmail.com).

## License
[MIT](https://choosealicense.com/licenses/mit/)