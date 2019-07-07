# Golang Shopping Cart

## Installation

#### Clone the source

```bash
git clone https://github.com/taufikardiyan28/mataharitest1.git
```

#### Setup dependencies

```bash
go build
```

#### Configuration
Edit config.yaml file and change your port and mysql database configuration

#### Run the app
```
go run main.go
or
./mataharitest1
```

### Usage

##### Products:
| | path | method|data|
| ----- | ------ | ------ |-----|
|List All Products|/api/v1/products |GET||
    
##### Carts:
| | path | method|data|
| ----- | ------ | ------ |-----|
|List|/api/v1/carts |GET||
|Add Product to Cart|/api/v1/carts |POST|JSON{product_id:(int), qty: (int)}|
|Update Cart|/api/v1/carts/:id |PUT|JSON{product_id:(int), qty: (int)}|
|Delete Cart|/api/v1/carts/:id |DELETE||
