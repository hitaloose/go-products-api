# Go Products API

A modern and lightweight API for managing product data built with the **Gin framework** in Go. This project features clean architecture, modular design, and mock data handling for testing and development.

---

## Features

- **Product Management**: Create new product records with essential details like name, description, price, stock, and category.
- **Lightweight and Fast**: Built with the Gin framework, providing high performance.
- **Mocked Repository**: Simulates database operations for seamless development and testing.
- **Clean Architecture**: Separation of concerns with services, repositories, and controllers.
- **Custom Error Handling**: Graceful handling of errors with descriptive HTTP responses.
- **Validation and Binding**: JSON request body validation and error formatting.

---

## Project Structure

```
.
├── cmd/               # Entry point of the application
│   └── main.go  # Initializes the Gin server and routes
├── scripts/           # Scripts for running and development
└── internal/           # Core business logic
    ├── controllers/    # Request handlers
    ├── dtos/           # Data Transfer Objects
    ├── exceptions/     # Custom error types and handlers
    ├── helpers/        # Utility functions
    ├── models/         # Data models
    ├── repositories/   # Data access layer
    ├── router/         # API routes
    └── services/       # Business logic
```

---

## Requirements

- **Go**: Version 1.23.3 or later
- **Gin Framework**: Used for HTTP handling

Dependencies are managed via Go modules and listed in `go.mod`.

---

## Getting Started

### Clone the Repository

```bash
$ git clone https://github.com/hitaloose/go-products-api.git
$ cd go-products-api
```

### Install Dependencies

```bash
$ go mod tidy
```

### Run the Development Server

```bash
$ ./scripts/dev.sh
```

The API will be accessible at `http://localhost:8080`.

---

## API Endpoints

### Create a Product

**POST** `/product`

**Request Body:**

```json
{
  "name": "Sample Product",
  "description": "A description of the product",
  "price": 99.99,
  "stock": 10,
  "category": "Electronics"
}
```

**Response:**

- **201 Created**

```json
{
  "product": {
    "id": 1,
    "name": "Sample Product",
    "description": "A description of the product",
    "price": 99.99,
    "stock": 10,
    "category": "Electronics",
    "created_at": "2024-12-08T12:00:00Z",
    "updated_at": "2024-12-08T12:00:00Z"
  }
}
```

**Error Responses:**

- **400 Bad Request**: Invalid input data.
- **422 Unprocessable Entity**: Validation errors.

---

## Development

### Run in Debug Mode

```bash
$ ./scripts/dev.sh
```

### Run in Production Mode

```bash
$ ./scripts/run.sh
```

---

## Technologies Used

- **[Gin](https://gin-gonic.com/)**: High-performance HTTP web framework.
- **[Go Playground Validator](https://github.com/go-playground/validator)**: Request validation library.
- **Go Modules**: Dependency management.

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a clear description of your changes.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

## Acknowledgements

- Inspired by modern Go backend practices and clean architecture principles.
- Special thanks to the developers behind Gin and other libraries used in this project.
