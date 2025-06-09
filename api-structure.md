# Posts API - Project Structure

This document outlines the project structure for the Go Posts API using the standard library.

## Directory Structure

```
posts-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   ├── database.go
│   │   └── env.go
│   ├── models/
│   │   ├── post.go
│   │   └── user.go                    # Reference model for Users API
│   ├── repository/
│   │   └── post_repository.go
│   ├── services/
│   │   ├── post_service.go
│   │   └── user_service.go            # HTTP client to Users API
│   ├── handlers/
│   │   └── post_handler.go
│   ├── middleware/
│   │   ├── auth.go                    # JWT validation
│   │   └── cors.go
│   └── dto/
│       ├── request/
│       │   ├── create_post.go
│       │   └── update_post.go
│       └── response/
│           └── post_response.go
├── pkg/
│   └── utils/
│       └── response.go                # Standardized responses
├── go.mod
├── go.sum
├── .env
└── README.md
```

## Directory Descriptions

### `cmd/`

Application entry points. Contains the main function and application bootstrap code.

- **`server/main.go`**: Main application entry point, server initialization

### `internal/`

Private application code that cannot be imported by other projects.

#### `config/`

Configuration management and setup.

- **`database.go`**: Database connection and configuration
- **`env.go`**: Environment variables loading and validation

#### `models/`

Data models and entity definitions.

- **`post.go`**: Post entity model for database operations
- **`user.go`**: User reference model for external Users API integration

#### `repository/`

Data access layer with database operations.

- **`post_repository.go`**: Post data access methods and database queries

#### `services/`

Business logic layer containing core application functionality.

- **`post_service.go`**: Post business logic and operations
- **`user_service.go`**: HTTP client for external Users API communication

#### `handlers/`

HTTP request handlers (controllers) for API endpoints.

- **`post_handler.go`**: HTTP handlers for post-related endpoints

#### `middleware/`

HTTP middleware for request processing.

- **`auth.go`**: Authentication middleware for JWT validation via Users API
- **`cors.go`**: CORS middleware for cross-origin request handling

#### `dto/`

Data Transfer Objects for API request and response structures.

##### `request/`

- **`create_post.go`**: Request DTO for post creation
- **`update_post.go`**: Request DTO for post updates

##### `response/`

- **`post_response.go`**: Response DTO for post data

### `pkg/`

Public library code that can be imported by other projects.

#### `utils/`

- **`response.go`**: Standardized API response utilities and helpers

### Configuration Files

- **`go.mod`**: Go module definition and dependencies
- **`go.sum`**: Dependency checksums
- **`.env`**: Environment variables (not tracked in git)
- **`README.md`**: Project documentation

## Architecture Principles

This structure follows:

- **Clean Architecture**: Clear separation of concerns with distinct layers
- **Go Conventions**: Standard Go project layout and naming conventions
- **Microservice Pattern**: External authentication delegation to Users API
- **Standard Library First**: Built using Go's standard `net/http` library
