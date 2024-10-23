# GoTu Bookstore

An API for online bookstore in Go.

## Installing

```
git clone git@github.com:rudysuharyadi/gotu-bookstore.git
cd gotu-bookstore
go mod download
```

## Docker

To run this application, you need PostgreSQL and Redis. You can simply just run docker compose.

```
docker compose up
```

## Setup

Setup can be found in the application.json inside resources folder.

```
cmd/gotu-bookstore/resource/application.json
```

### Server

| Server Setup    | Value         |
| --------------- | ------------- |
| SERVER_APP_MODE | "development" |
| SERVER_APP_HOST | "127.0.0.1"   |
| SERVER_APP_PORT | 4041          |

### Database

| Database Setup             | Value       |
| -------------------------- | ----------- |
| DATABASE_DB_NAME           | gotu        |
| DATABASE_HOST              | "127.0.0.1" |
| DATABASE_PORT              | 5432        |
| DATABASE_USERNAME          | "postgres"  |
| DATABASE_PASSWORD          | "postgres"  |
| DATABASE_MAX_IDDLE_CONN    | 5           |
| DATABASE_MAX_OPEN_CONN     | 10          |
| DATABASE_CONN_MAX_LIFETIME | 1           |
| DATABASE_LOG_MODE          | 1           |

### Redis

| Redis Setup                    | Value       |
| ------------------------------ | ----------- |
| REDIS_HOST                     | "127.0.0.1" |
| REDIS_PORT                     | 6379        |
| REDIS_DEFAULT_CACHE_EXPIRATION | 2592000     |
| REDIS_USERNAME                 | "default"   |
| REDIS_PASSWORD                 | "redis"     |

### Auth

| Auth Setup                    | Value    |
| ----------------------------- | -------- |
| AUTH_ACCESS_TOKEN_SECRET_KEY  | "secret" |
| AUTH_ACCESS_TOKEN_EXPIRATION  | 604800   |
| AUTH_REFRESH_TOKEN_SECRET_KEY | "secret  |
| AUTH_REFRESH_TOKEN_EXPIRATION | 2629743  |
| AUTH_ENCRYPTION_KEY           | "secret" |

For authentication key, we can use existing key in the application.json. You can generate a new one and replace it.

Generate a Private Key:
Use the openssl ecparam command to generate a private key with the P-256 curve (also known as prime256v1), which corresponds to ES256:

```
openssl ecparam -name prime256v1 -genkey -noout -out private_key.pem
```

This command generates a private key and saves it in the private_key.pem file.

Extract the Public Key from the Private Key:
You can extract the public key from the private key using the openssl ec command:

```
openssl ec -in private_key.pem -pubout -out public_key.pem
```

This command extracts the public key and saves it in the public_key.pem file.

Command to generate a 256-bit (32-byte) random key:

```
openssl rand -hex 32
```

This command generates 32 random bytes and then converts them to a hexadecimal string representation. Each byte consists of two hexadecimal characters, resulting in a 64-character string (32 bytes \* 2 characters/byte = 64 characters).

## Run Server

Because resources folder is under cmd/gotu-bookstore, you need to execute run command inside cmd/gotu-bookstore folder.

```
cd cmd/gotu-bookstore
go run .
```

In VSCode, you can set launch.json to be something like this

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/gotu-bookstore"
        }
    ]
}
```

## Test

To run the test, simply just execute command:

```
go test ./...
```

To create a mock object, use mockery command in root folder.

```
mockery
```

Make sure you run mockery command in root folder. The mock object created will be placed under mocks folder.

## API Specs

You can found the api-spec.json in root folder. There is also postman collection.

### Account Service

| API Path                          | Method | What it does  |
| --------------------------------- | ------ | ------------- |
| /account-service/v1/logout        | POST   | Logout        |
| /account-service/v1/login         | POST   | Login         |
| /account-service/v1/register      | POST   | Register      |
| /account-service/v1/refresh-token | POST   | Refresh Token |

### Product Service

| API Path                           | Method | What it does                          |
| ---------------------------------- | ------ | ------------------------------------- |
| /product-service/v1/books          | GET    | Get all books, paginated, with search |
| /product-service/v1/books/:book_id | GET    | Get book details                      |

### Order Service

| API Path                                       | Method | What it does                                 |
| ---------------------------------------------- | ------ | -------------------------------------------- |
| /order-service/v1/shopping-cart                | GET    | Get current shopping cart from existing user |
| /order-service/v1/shopping-cart                | POST   | Add / Edit / Delete item on shopping cart    |
| /order-service/v1/shopping-cart/clear          | POST   | Clear current shopping cart                  |
| /order-service/v1/shopping-cart/checkout       | POST   | Checkout                                     |
| /order-service/v1/transactions                 | GET    | Get history transaction from existing user   |
| /order-service/v1/transactions/:transaction_id | GET    | Get transaction details                      |
