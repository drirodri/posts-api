# Posts API - Go Microservice for Post Management

A high-performance RESTful API for managing posts built with Go using the standard library. This service is part of the Dri Posts microservices architecture, focusing on post CRUD operations with delegated authentication to the Users API.

## 🏗️ Architecture Overview

This Posts API is part of a microservices architecture:

- **[Frontend](https://github.com/drirodri/posts-frontend)** (React + TypeScript + Vite) - User interface
- **[Users API](https://github.com/drirodri/users-api)** (NestJS) - Authentication & user management
- **Posts API** (Go + Standard Library) - This service for posts management

## 🚀 Features

### ✅ Implemented Features

- 📝 **Post Management** - Complete CRUD operations for posts
- 🔒 **Authentication Integration** - JWT validation via Users API
- 🛡️ **Authorization** - Role-based access control (author-only operations)
- 🌐 **CORS Support** - Cross-origin resource sharing configuration
- 📊 **Standardized Responses** - Consistent API response format
- 🏗️ **Clean Architecture** - Separation of concerns with layered structure
- 📚 **Bruno API Collection** - Complete testing suite with Bruno

### 🚧 Planned Features

- 🔍 **Advanced Search** - Full-text search and filtering capabilities
- 📄 **Pagination** - Efficient data pagination for large datasets
- 🏷️ **Tags & Categories** - Post categorization and tagging system
- 📊 **Analytics** - Post engagement and statistics tracking
- 🔄 **Caching Layer** - Redis integration for improved performance
- 📱 **Image Upload** - Media attachment support for posts

## 🛠️ Technology Stack

### Core Technologies

- **Go 1.21+** - High-performance programming language
- **net/http** - Go standard library HTTP server
- **Gorilla Mux** - HTTP router for advanced routing patterns
- **GORM** - Feature-rich ORM for Go
- **PostgreSQL** - Robust relational database

### Development Tools

- **godotenv** - Environment variable management
- **Bruno** - API testing and documentation
- **Air** - Live reload for Go development (optional)
- **Go Modules** - Dependency management

### Architecture Patterns

- **Clean Architecture** - Clear separation of concerns
- **Repository Pattern** - Data access abstraction
- **Service Layer Pattern** - Business logic encapsulation
- **Microservices Pattern** - Distributed system architecture
- **External Authentication** - Delegated auth to Users API

## 🏁 Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- Running [Users API](https://github.com/drirodri/users-api) (for authentication)
- [Bruno API Client](https://www.usebruno.com/) (for testing)

### Installation

1. **Navigate to posts-api directory:**

   ```bash
   cd posts-api
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Set up environment variables:**

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Configure database connection:**

   ```env
   # Database Configuration
   DATABASE_HOST=localhost
   DATABASE_PORT=5432
   DATABASE_USERNAME=postgres
   DATABASE_PASSWORD=your_password
   DATABASE_NAME=posts_api_db

   # Application Configuration
   PORT=8080

   # Users API Integration
   USERS_API_URL=http://localhost:3000
   ```

5. **Start the server:**

   ```bash
   go run cmd/server/main.go
   ```

6. **Verify installation:**
   ```
   http://localhost:8080/health
   ```

### Available Commands

```bash
# Run development server
go run cmd/server/main.go

# Build production binary
go build -o posts-api cmd/server/main.go

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...

# Vet code for issues
go vet ./...

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy
```

## 📁 Project Structure

```
posts-api/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   ├── database.go          # Database configuration
│   │   └── env.go               # Environment variables
│   ├── models/
│   │   ├── post.go              # Post entity model
│   │   └── user.go              # User reference model
│   ├── repository/
│   │   └── post_repository.go   # Data access layer
│   ├── services/
│   │   ├── post_service.go      # Business logic layer
│   │   └── user_service.go      # Users API client
│   ├── handlers/
│   │   └── post_handler.go      # HTTP request handlers
│   ├── middleware/
│   │   ├── auth.go              # JWT authentication
│   │   └── cors.go              # CORS configuration
│   └── dto/
│       ├── request/
│       │   ├── create_post.go   # Post creation DTO
│       │   └── update_post.go   # Post update DTO
│       └── response/
│           └── post_response.go # Post response DTO
├── pkg/
│   └── utils/
│       └── response.go          # Standardized responses
├── posts-bruno-api-requests/    # Bruno API testing collection
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── .env.example                 # Environment template
├── .gitignore                   # Git ignore rules
├── api-structure.md             # Project structure documentation
├── go-conventions.md            # Go coding conventions
├── GUIDE.md                     # Development guide
└── README.md                    # This file
```

## 🔌 API Integration

### Users API Integration

- **Authentication Validation** - JWT token verification
- **User Information** - User details retrieval
- **Role Verification** - Permission checking
- **Base URL**: Configure in environment variables

### API Endpoints

#### Public Endpoints

```
GET    /health              # Health check
GET    /                    # API information
GET    /api/v1/posts        # Get all posts
GET    /api/v1/posts/:id    # Get post by ID
GET    /api/v1/posts/author/:authorId # Get posts by author
```

#### Protected Endpoints (Require JWT)

```
POST   /api/v1/posts        # Create new post
PUT    /api/v1/posts/:id    # Update post (author only)
DELETE /api/v1/posts/:id    # Delete post (author only)
```

### Authentication Flow

1. **Client authenticates** with Users API to get JWT token
2. **Client includes token** in Authorization header: `Bearer <token>`
3. **Posts API validates** token by calling Users API validation endpoint
4. **If valid**, request proceeds; **if invalid**, returns 401 Unauthorized

### Request/Response Format

#### Create Post Request

```json
{
  "title": "My Amazing Post",
  "content": "This is the content of my post with detailed information."
}
```

#### Standard Response Format

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {
    "id": 1,
    "title": "My Amazing Post",
    "content": "This is the content...",
    "author_id": 123,
    "created_at": "2025-06-11T10:00:00Z",
    "updated_at": "2025-06-11T10:00:00Z"
  }
}
```

#### Error Response Format

```json
{
  "success": false,
  "message": "Error description",
  "error": {
    "code": "ERROR_CODE",
    "message": "Detailed error message"
  }
}
```

## 🧪 API Testing

### Bruno Collection

Complete API testing suite available in `posts-bruno-api-requests/`:

1. **Import Collection** - Open Bruno and import the collection
2. **Configure Environment** - Select "Local" environment
3. **Get JWT Token** - Login to Users API first
4. **Run Tests** - Execute requests in sequence

### Quick Test Sequence

```bash
# 1. Health Check
GET /health

# 2. Get All Posts (should be empty initially)
GET /api/v1/posts

# 3. Create Post (requires JWT token)
POST /api/v1/posts

# 4. Get Created Post
GET /api/v1/posts/{id}

# 5. Update Post (author only)
PUT /api/v1/posts/{id}

# 6. Delete Post (author only)
DELETE /api/v1/posts/{id}
```

### Testing Resources

- **[Quick Start Guide](posts-bruno-api-requests/QUICK_START.md)** - 5-minute setup
- **[Complete Collection](posts-bruno-api-requests/README.md)** - Full testing documentation
- **[API Structure](api-structure.md)** - Project architecture details
- **[Development Guide](GUIDE.md)** - Comprehensive development guide

## 🧪 Development Guidelines

### Code Style

- **Go Conventions** - Follow standard Go naming and structure
- **gofmt** - Automatic code formatting
- **golint** - Code quality checking
- **go vet** - Static analysis for bugs

### Architecture Guidelines

- **Clean Architecture** - Maintain clear separation of concerns
- **Repository Pattern** - Abstract data access logic
- **Service Layer** - Encapsulate business logic
- **Handler Layer** - Handle HTTP requests/responses
- **Middleware** - Cross-cutting concerns (auth, CORS, logging)

### Error Handling

- **Explicit Error Returns** - Go's explicit error handling pattern
- **Standardized Responses** - Consistent error response format
- **HTTP Status Codes** - Proper HTTP status code usage
- **Validation Errors** - Clear validation error messages

## 🔄 Integration with Backend Services

### Users API (NestJS)

- **JWT Validation** - Token verification endpoint calls
- **User Information** - User details retrieval
- **Role Management** - Permission and role checking
- **Security** - Secure token validation process

### Database (PostgreSQL)

- **GORM ORM** - Database operations and migrations
- **Connection Pool** - Efficient database connections
- **Transactions** - ACID compliance for data operations
- **Indexing** - Optimized queries for performance

## 📊 Performance & Monitoring

### Performance Features

- **Go Concurrency** - Goroutines for concurrent request handling
- **Connection Pooling** - Efficient database connections
- **Standard Library** - Lightweight HTTP server implementation
- **Minimal Dependencies** - Fast startup and low memory footprint

### Monitoring (Planned)

- **Health Endpoints** - Service health monitoring
- **Metrics Collection** - Performance metrics tracking
- **Logging** - Structured logging for debugging
- **Distributed Tracing** - Request tracing across services

## 📚 Documentation

### API Documentation

- **Bruno Collection** - Interactive API testing and documentation
- **API Structure** - Detailed project structure documentation
- **Development Guide** - Step-by-step implementation guide
- **Go Conventions** - Go-specific coding standards

### Learning Resources

- [Go Documentation](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [GORM Documentation](https://gorm.io/docs/)
- [Gorilla Mux Documentation](https://pkg.go.dev/github.com/gorilla/mux)

## 🤝 Contributing

1. Follow Go conventions and coding standards
2. Write meaningful commit messages following conventional commits
3. Include tests for all new features
4. Update documentation for any API changes
5. Ensure all tests pass before committing
6. Use `gofmt` and `go vet` before submitting

## 🚀 Deployment

### Build Process

```bash
# Build production binary
go build -o posts-api cmd/server/main.go

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o posts-api-linux cmd/server/main.go
```

### Deployment Options

- **Docker** - Containerized deployment
- **Binary Deployment** - Direct binary execution
- **Cloud Platforms** - AWS, GCP, Azure deployment
- **Kubernetes** - Container orchestration

## 📝 License

This project is for educational purposes and demonstrates modern Go development practices with microservices architecture, clean code principles, and RESTful API design.

---

**Note**: This is part of a comprehensive full-stack application showcasing modern web development practices with React frontend, NestJS authentication service, and Go posts API. The implementation demonstrates production-ready patterns for building scalable microservices architectures.
