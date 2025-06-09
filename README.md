# Posts API

A RESTful API for managing posts built with Go using the standard library.

## Features

- Post management (CRUD operations)
- Authentication delegation to [Users API](https://github.com/drirodri/users-api)
- Integration with external Users API for user validation
- CORS support
- Standardized API responses
- Built with Go's standard `net/http` library

## Architecture

This API follows a microservice architecture pattern where:

- **Posts API** (this service): Handles post-related operations
- **Users API**: Handles authentication, user management, and JWT validation
- Authentication tokens are validated against the Users API for each protected request

## Project Structure

See `api-structure.md` for detailed project structure.

## Getting Started

1. Install dependencies:

   ```bash
   go mod download
   ```

2. Set up environment variables in `.env`:

   ```bash
   # Database Configuration
   DATABASE_HOST=localhost
   DATABASE_PORT=your_db_port
   DATABASE_USERNAME=postgres
   DATABASE_PASSWORD=your_password
   DATABASE_NAME=posts-api

   # Application Configuration
   PORT=your_desired_api_port

   # Users API Integration
   USERS_API_URL=https://your-users-api.com
   ```

3. Ensure the Users API is running and accessible

4. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Public Endpoints

- `GET /health` - Health check

### Protected Endpoints (require authentication via Users API)

- `GET /posts` - Get all posts
- `GET /posts/:id` - Get post by ID
- `POST /posts` - Create new post (requires valid user token)
- `PUT /posts/:id` - Update post (requires ownership or admin role)
- `DELETE /posts/:id` - Delete post (requires ownership or admin role)

## Authentication

This API uses token-based authentication managed by the [Users API](https://github.com/drirodri/users-api).

### How it works:

1. Client authenticates with Users API to get a JWT token
2. Client includes token in Authorization header: `Bearer <token>`
3. Posts API validates token by calling Users API validation endpoint
4. If valid, request proceeds; if invalid, returns 401 Unauthorized

### Request Headers

```
Authorization: Bearer <jwt_token_from_users_api>
Content-Type: application/json
```

## Technology Stack

- **Language**: Go (Golang)
- **HTTP Framework**: Go standard library (`net/http`)
- **Router**: Gorilla Mux (optional)
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: Delegated to Users API
- **Environment**: godotenv

## Dependencies

- `github.com/gorilla/mux` - HTTP router
- `gorm.io/gorm` - ORM for database operations
- `gorm.io/driver/postgres` - PostgreSQL driver for GORM
- `github.com/joho/godotenv` - Environment variable loader
