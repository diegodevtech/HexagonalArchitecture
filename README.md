# Hexagonal Architecture (Ports and Adapters) Project

<img width="400" height="400" src="https://solutions.lykdat.com/blog/content/images/2023/08/image-4.png" alt="Exemplo imagem">

> Main approach: Decouples the application core from external libraries, bringing them only as adapters. Focus on scalability.

## 💻 Prerequisites

Before we get started, Take a look if you have the following:

- `docker v27.3.1` or higher

## 🚀 Installing

First of all, start with the `git clone` command.
Then,

Open-up your terminal and run:
```
docker compose up -d
```

## ☕ Using
##### OBS.: You can use the given `db.sqlite` file to show up some already created products.

This project contains a core application and the following adapters:
- CLI
- Web Server
- Databse (Sqlite3)

### You can meke some requests using the CLI

##### Getting all products: 
```
go run main.go cli -a=all
```

##### Getting product by ID: 
```
go run main.go cli -a=get -i={id}
```

##### Creating product: 
```
go run main.go cli -a=create -n="{Product'a Name}" -p={some float64 value}
```

### You can also start using the Webserver running:
```
go run main.go http
```
##### OBS.: Choose the HTTP CLient of  your preference ant start making requests, such as:
##### Getting all products: 
```
  http://localhost:9000/product
```
##### Getting product by Id: 
```
  GET - http://localhost:9000/product/{productId}
```
##### Getting all products: 
```
  GET - http://localhost:9000/product
```
##### Creating a product: 
```
  POST - http://localhost:9000/product
  BODY:
  {
    "name": "ThunderClient Product",
    "price": "100.15"
  }
```
