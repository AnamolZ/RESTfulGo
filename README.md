# RESTfulGo

RESTfulGo is a lightweight and easy-to-use RESTful API service developed in Go (Golang). This project serves as a boilerplate for building scalable web applications and APIs with Go, providing a solid foundation for future development.

## Features

- Simple and clean RESTful API design
- Easy to extend and modify
- Built-in routing using the Go standard library
- Support for JSON input/output
- Error handling and logging

Requirements
```bash
go mod init .
go get github.com/julienschmidt/httprouter
```

1. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

### API Endpoints

- `GET /api/example` - Retrieves data resource.
- `POST /api/example` - Creates a new data resource.
- `PUT /api/example/{id}` - Updates an existing data resource.
- `DELETE /api/example/{id}` - Deletes an data resource.

## Author

[Anamol Dhakal](https://www.anamoldhakal.com.np)