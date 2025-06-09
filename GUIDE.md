# Go Posts API Development Guide

## Overview

This guide will help you build a RESTful Posts API using Go. Since you have experience with Java Spring Boot and NestJS, I'll draw parallels to help you understand Go's approach to building web APIs.

## Table of Contents

1. [Initial Setup](#initial-setup)
2. [Project Structure Overview](#project-structure-overview)
3. [Development Workflow](#development-workflow)
4. [File-by-File Implementation Guide](#file-by-file-implementation-guide)
5. [Git Workflow](#git-workflow)
6. [Testing Strategy](#testing-strategy)
7. [Deployment Considerations](#deployment-considerations)

---

## Initial Setup

### Step 1: Initialize Go Module

First, initialize your Go module (similar to `package.json` in Node.js or `pom.xml` in Maven):

```bash
go mod init posts-api
```

**Go Documentation**: [Go Modules](https://go.dev/doc/modules/gomod-ref)

### Step 2: Install Dependencies

You'll need these essential packages:

- **net/http** (Go's standard HTTP library - no external web framework needed)
- **Gorilla Mux** (HTTP router for more advanced routing - similar to Express Router)
- **GORM** (ORM - similar to TypeORM or JPA)
- **Godotenv** (environment variables)
- **PostgreSQL Driver** (or your preferred database)

**Note**: We're using Go's standard `net/http` package instead of Gin for a more lightweight approach. User authentication and JWT handling is managed by the separate Users API at https://github.com/drirodri/users-api

**Go Documentation**: [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)

---

## Project Structure Overview

Our structure follows Go conventions and clean architecture principles (similar to what you'd see in well-structured Spring Boot applications):

- **`cmd/`**: Entry points (like your `main` class in Spring Boot)
- **`internal/`**: Private application code (not importable by other projects)
- **`pkg/`**: Public library code (reusable utilities)
- **`config/`**: Configuration management (like `@Configuration` classes)
- **`models/`**: Data entities (like JPA entities or TypeORM entities)
- **`repository/`**: Data access layer (like Spring Data repositories)
- **`services/`**: Business logic (like Spring `@Service` classes)
- **`handlers/`**: HTTP controllers (like Spring `@RestController` or NestJS controllers)
- **`middleware/`**: HTTP middleware (like Spring interceptors or NestJS guards)
- **`dto/`**: Data Transfer Objects (like Spring DTOs or NestJS DTOs)

**Go Documentation**: [Project Layout](https://github.com/golang-standards/project-layout)

---

## Development Workflow

### Phase 1: Foundation (Days 1-2)

1. Set up basic project structure
2. Configure environment and database
3. Create basic models
4. Set up main server entry point

### Phase 2: Core Business Logic (Days 3-4)

1. Implement repository layer
2. Build service layer
3. Create DTOs for request/response handling

### Phase 3: HTTP Layer (Days 5-6)

1. Implement handlers/controllers
2. Set up routing
3. Add middleware (auth, CORS)

### Phase 4: Integration & Testing (Days 7-8)

1. Integrate with external Users API
2. Add comprehensive testing
3. Error handling and validation

---

## File-by-File Implementation Guide

### 1. Configuration Layer

#### `internal/config/env.go`

**Purpose**: Environment variable configuration (like Spring's `@ConfigurationProperties`)
**What to implement**:

- Load environment variables from `.env` file
- Define configuration struct with all app settings
- Validation for required environment variables
- Default values for optional settings

**Concepts**: Similar to Spring Boot's application.properties configuration
**Go Documentation**: [Environment Variables](https://pkg.go.dev/os#Getenv)

#### `internal/config/database.go`

**Purpose**: Database connection and configuration (like Spring's DataSource configuration)
**What to implement**:

- Database connection setup using GORM
- Connection pool configuration
- Database migration handling
- Health check functionality

**Concepts**: Similar to Spring Boot's DataSource and JPA configuration
**Go Documentation**: [Database/SQL](https://pkg.go.dev/database/sql)

### 2. Data Layer

#### `internal/models/post.go`

**Purpose**: Post entity definition (like JPA entities or TypeORM entities)
**What to implement**:

- Post struct with GORM tags
- Field validations
- Relationships (if any)
- JSON serialization tags

**Concepts**: Similar to `@Entity` classes in Spring or entity classes in TypeORM
**Go Documentation**: [Structs](https://go.dev/tour/moretypes/2)

#### `internal/models/user.go`

**Purpose**: User reference model for external API integration
**What to implement**:

- User struct representing external API response
- Only fields needed for post operations
- JSON deserialization tags

**Concepts**: Similar to DTOs for external service responses
**Go Documentation**: [JSON and Go](https://go.dev/blog/json)

#### `internal/repository/post_repository.go`

**Purpose**: Data access layer (like Spring Data repositories)
**What to implement**:

- Interface definition for repository contract
- GORM-based implementation
- CRUD operations (Create, Read, Update, Delete)
- Query methods with filtering and pagination
- Error handling for database operations

**Concepts**: Similar to Spring Data JPA repositories or TypeORM repositories
**Go Documentation**: [Interfaces](https://go.dev/tour/methods/9)

### 3. Business Logic Layer

#### `internal/services/post_service.go`

**Purpose**: Business logic for post operations (like Spring `@Service` classes)
**What to implement**:

- Interface definition for service contract
- Business rules and validation
- Orchestration between repository and external services
- Transaction management
- Business-level error handling

**Concepts**: Similar to Spring Service layer or NestJS Service providers
**Go Documentation**: [Error Handling](https://go.dev/blog/error-handling-and-go)

#### `internal/services/user_service.go`

**Purpose**: HTTP client for Users API integration (like Spring's RestTemplate or NestJS HttpService)
**What to implement**:

- HTTP client setup
- User validation methods
- API request/response handling
- Retry logic and error handling
- Caching (if needed)

**Concepts**: Similar to Feign clients in Spring or HttpService in NestJS
**Go Documentation**: [HTTP Client](https://pkg.go.dev/net/http#Client)

### 4. Data Transfer Objects

#### `internal/dto/create_post.go`

**Purpose**: Request DTO for creating posts (like Spring DTOs with validation annotations)
**What to implement**:

- Struct for incoming POST request data
- Validation tags for required fields
- JSON binding tags
- Transformation methods to domain models

**Concepts**: Similar to Spring's request DTOs with `@Valid` annotations
**Go Documentation**: [Struct Tags](https://go.dev/wiki/Well-known-struct-tags)

#### `internal/dto/update_post.go`

**Purpose**: Request DTO for updating posts
**What to implement**:

- Struct for PATCH/PUT request data
- Optional field handling
- Validation for update operations
- Partial update support

#### `internal/dto/post_response.go`

**Purpose**: Response DTO for post data (like Spring ResponseEntity bodies)
**What to implement**:

- Struct for API responses
- Pagination support
- Metadata inclusion
- Consistent response formatting

### 5. HTTP Layer

#### `internal/handlers/post_handler.go`

**Purpose**: HTTP controllers using Go's standard net/http (like Spring `@RestController` or NestJS controllers)
**What to implement**:

- Handler struct with service dependencies
- HTTP endpoint methods (GET, POST, PUT, DELETE)
- Request parsing and validation using standard library or Gorilla Schema
- Response formatting with JSON encoding
- HTTP status code handling
- Error response formatting
- Route registration with Gorilla Mux or standard ServeMux

**Concepts**: Similar to Spring Controllers or NestJS Controllers, but using Go's standard HTTP handlers
**Go Documentation**: [HTTP Handlers](https://pkg.go.dev/net/http#Handler)

#### `internal/middleware/auth.go`

**Purpose**: Authentication middleware for validating requests against Users API (like Spring Security or NestJS Guards)
**What to implement**:

- HTTP client to validate tokens with Users API (https://github.com/drirodri/users-api)
- Token extraction from Authorization header
- User context extraction from Users API response
- Authorization logic based on user roles/permissions
- Error handling for invalid tokens or failed API calls
- Caching of validation results (optional, for performance)

**Concepts**: Similar to Spring Security filters or NestJS AuthGuard, but delegating to external service
**Go Documentation**: [Middleware Patterns](https://go.dev/doc/articles/wiki/)

#### `internal/middleware/cors.go`

**Purpose**: CORS handling (like Spring's CORS configuration or NestJS CORS)
**What to implement**:

- CORS headers configuration
- Preflight request handling
- Origin validation
- Credentials handling

**Concepts**: Similar to Spring's `@CrossOrigin` or NestJS CORS configuration

### 6. Utilities

#### `pkg/utils/response.go`

**Purpose**: Standardized API responses (like Spring's ResponseEntity or custom response wrappers)
**What to implement**:

- Standard response structure
- Success response helpers
- Error response helpers
- Pagination response wrapper
- Consistent status code mapping

**Concepts**: Similar to standardized response wrappers in Spring or NestJS

### 7. Application Entry Point

#### `cmd/server/main.go`

**Purpose**: Application bootstrap using Go's standard HTTP server (like Spring Boot's main class with `@SpringBootApplication`)
**What to implement**:

- HTTP server initialization using `net/http`
- Dependency injection setup (manual in Go)
- Route registration with Gorilla Mux or standard ServeMux
- Middleware registration (auth, CORS, logging)
- Graceful shutdown handling with context cancellation
- Server configuration (port, timeouts, etc.)
- Integration with Users API for authentication

**Concepts**: Similar to Spring Boot's main method or NestJS bootstrap, but using standard library
**Go Documentation**: [HTTP Server](https://pkg.go.dev/net/http#Server)

---

## Git Workflow

### Branch Naming Convention

- **Feature branches**: `feature/TICKET-ID-short-description`
- **Bug fixes**: `bugfix/TICKET-ID-short-description`
- **Hotfixes**: `hotfix/TICKET-ID-short-description`
- **Releases**: `release/v1.0.0`

### Recommended Git Flow

#### 1. Feature Development

```
main
 └── develop
     └── feature/API-001-implement-post-crud
```

#### 2. Branch Strategy

- **`main`**: Production-ready code
- **`develop`**: Integration branch for features
- **Feature branches**: Individual features
- **Release branches**: Preparation for production releases

#### 3. Commit Message Convention

```
type(scope): description

feat(posts): implement post creation endpoint
fix(auth): resolve JWT token validation issue
docs(readme): update API documentation
test(posts): add unit tests for post service
```

#### 4. Pull Request Process

1. Create feature branch from `develop`
2. Implement feature with tests
3. Create PR to `develop`
4. Code review and approval
5. Merge to `develop`
6. Deploy to staging for testing
7. Create release branch when ready
8. Merge release to `main` and tag

### Development Phases with Branches

#### Phase 1: Foundation

- `feature/API-001-project-setup`
- `feature/API-002-database-config`
- `feature/API-003-basic-models`

#### Phase 2: Core Logic

- `feature/API-004-post-repository`
- `feature/API-005-post-service`
- `feature/API-006-dtos`

#### Phase 3: HTTP Layer

- `feature/API-007-post-handlers`
- `feature/API-008-middleware`
- `feature/API-009-routing`

#### Phase 4: Integration

- `feature/API-010-user-service-integration`
- `feature/API-011-error-handling`
- `feature/API-012-testing`

---

## Testing Strategy

### Unit Testing

- Test each layer independently
- Mock dependencies (similar to Mockito in Spring)
- Focus on business logic in services
- Test validation in DTOs

### Integration Testing

- Test HTTP endpoints end-to-end
- Test database operations
- Test external API integration

### Tools

- **Go's built-in testing**: `go test`
- **Testify**: Assertions and mocking
- **HTTP testing**: `httptest` package

**Go Documentation**: [Testing](https://go.dev/doc/tutorial/add-a-test)

---

## Key Differences from Spring Boot/NestJS

## Key Differences from Spring Boot/NestJS

### Web Framework Approach

- **Spring/NestJS**: Full-featured frameworks with extensive middleware ecosystems
- **Go**: Using standard library `net/http` with optional routing libraries like Gorilla Mux

### Authentication Strategy

- **Spring/NestJS**: Built-in authentication with JWT libraries
- **Go Posts API**: Delegates authentication to external Users API for microservice architecture

### Dependency Injection

- **Spring/NestJS**: Automatic DI container
- **Go**: Manual dependency injection in main.go

### Error Handling

- **Spring/NestJS**: Exception-based
- **Go**: Explicit error returns

### Configuration

- **Spring**: Annotations and auto-configuration
- **Go**: Explicit configuration in code

### Middleware/Interceptors

- **Spring**: AOP and interceptors
- **NestJS**: Decorators and guards
- **Go**: Function-based middleware chain with standard HTTP handlers

---

## Learning Resources

### Essential Go Documentation

- [A Tour of Go](https://go.dev/tour/) - Interactive tutorial
- [Effective Go](https://go.dev/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com/) - Code examples
- [Go Web Examples](https://gowebexamples.com/) - Web development patterns

### Framework Documentation

- [Gorilla Mux](https://pkg.go.dev/github.com/gorilla/mux) - HTTP router (if using)
- [GORM](https://gorm.io/docs/) - ORM documentation
- [net/http](https://pkg.go.dev/net/http) - Go standard HTTP package
- [Users API](https://github.com/drirodri/users-api) - External authentication service

---

## Next Steps

1. **Start with Phase 1**: Set up the foundation
2. **Follow the file-by-file guide**: Implement each component systematically
3. **Test as you go**: Write tests for each component
4. **Use Git flow**: Create branches for each feature
5. **Review and refactor**: Apply Go best practices

Remember: Go's philosophy is simplicity and explicitness. Unlike Spring Boot's "magic" annotations or NestJS decorators, Go requires you to be explicit about dependencies and configuration. This leads to more predictable and debuggable code.

Good luck with your Go journey! 🚀
