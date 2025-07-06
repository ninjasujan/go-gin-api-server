# Go Gin API Server

A RESTful API server built with Go and Gin framework, featuring a simple bookstore management system.

## 🚀 Features

- **RESTful API** with CRUD operations
- **Book management** (Create, Read, Delete)
- **JSON responses** with proper HTTP status codes
- **Built with Gin** - Fast HTTP web framework

## 📋 Prerequisites

- Go 1.24.2 or higher
- Git

## 🛠️ Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/go-gin-api-server.git
   cd go-gin-api-server
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Endpoints

### Books

| Method   | Endpoint     | Description         |
| -------- | ------------ | ------------------- |
| `GET`    | `/books`     | Get all books       |
| `POST`   | `/books`     | Create a new book   |
| `DELETE` | `/books/:id` | Delete a book by ID |

### Example Usage

**Get all books:**

```bash
curl http://localhost:8080/books
```

**Create a book:**

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"New Book","author":"New Author","price":150.00}'
```

**Delete a book:**

```bash
curl -X DELETE http://localhost:8080/books/1
```

## 📁 Project Structure

```
go-gin-api-server/
├── main.go          # Server entry point
├── routes/
│   └── routes.go    # API route definitions
├── go.mod           # Go module file
└── README.md        # This file
```

## 🛠️ Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **JSON** - Data format

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

**Happy coding! 🎉**
