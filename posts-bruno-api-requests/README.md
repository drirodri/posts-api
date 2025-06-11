# Posts API - Bruno Collection

This Bruno collection provides comprehensive API testing for the Posts API microservice.

## 📋 Collection Overview

This collection includes:

- **Health & Status Checks** - Basic API connectivity tests
- **Posts Management** - Complete CRUD operations for posts
- **Error Cases** - Negative testing scenarios
- **Test Workflows** - Sequential test scenarios

## 🚀 Getting Started

### Prerequisites

1. **Bruno API Client** - Download from [usebruno.com](https://www.usebruno.com/)
2. **Posts API** - Running on localhost:3334 (or update environment)
3. **Users API** - Running on localhost:3333 for authentication
4. **Valid JWT Token** - Obtain from Users API for protected endpoints

### Setup Instructions

1. **Import Collection**

   ```
   File -> Open Collection -> Select this folder
   ```

2. **Configure Environment**

   - Select "Local" environment for development
   - Update URLs in environments if needed:
     - `baseUrl`: http://localhost:3334
     - `usersApiUrl`: http://localhost:3333

3. **Obtain JWT Token**
   - Login to Users API to get a JWT token
   - Update `your-jwt-token-here` in protected requests

## 📁 Collection Structure

```
posts-bruno-api-requests/
├── bruno.json                    # Collection configuration
├── README.md                     # This file
├── environments/
│   ├── Local.bru                 # Local development environment
│   └── Production.bru            # Production environment
├── Health Check.bru              # API health verification
├── Root Endpoint.bru             # Basic API info
├── Posts/
│   ├── Get All Posts.bru         # List all posts (public)
│   ├── Get Post by ID.bru        # Get specific post (public)
│   ├── Get Posts by Author.bru   # Get posts by author (public)
│   ├── Create Post.bru           # Create new post (protected)
│   ├── Update Post.bru           # Update existing post (protected)
│   └── Delete Post.bru           # Delete post (protected)
├── Error Cases/
│   ├── Create Post - Invalid Token.bru
│   ├── Create Post - Missing Token.bru
│   ├── Create Post - Validation Errors.bru
│   └── Get Post - Invalid ID.bru
└── Test Workflows/
    └── Complete CRUD Workflow.bru
```

## 🔐 Authentication

The Posts API uses JWT tokens for authentication from the Users API.

### Getting a JWT Token

1. **Register/Login to Users API:**

   ```bash
   POST http://localhost:3333/auth/login
   Content-Type: application/json

   {
     "email": "user@example.com",
     "password": "your-password"
   }
   ```

2. **Extract the JWT token from the response**

3. **Update Bearer tokens in protected requests:**
   - Create Post
   - Update Post
   - Delete Post

### Token Usage

Replace `your-jwt-token-here` in the `auth:bearer` sections with your actual JWT token:

```
auth:bearer {
  token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
}
```

## 🧪 Test Scenarios

### 1. Basic Connectivity

- Health Check
- Root Endpoint

### 2. Public Endpoints (No Auth Required)

- Get All Posts
- Get Post by ID
- Get Posts by Author

### 3. Protected Endpoints (Auth Required)

- Create Post
- Update Post (only by author)
- Delete Post (only by author)

### 4. Error Handling

- Invalid/Missing tokens
- Validation errors
- Invalid IDs
- Permission denied scenarios

## 📝 Test Data Examples

### Valid Post Creation

```json
{
  "title": "My First Post",
  "content": "This is the content of my first post. It contains useful information about the topic."
}
```

### Post Update

```json
{
  "title": "Updated Post Title",
  "content": "This is the updated content with more details."
}
```

## ✅ Running Tests

### Sequential Testing (Recommended)

1. Start with Health Check
2. Run Get All Posts (should be empty initially)
3. Create a post (requires valid token)
4. Test read operations with the created post ID
5. Update the post
6. Test delete functionality

### Individual Endpoint Testing

- Each request can be run independently
- Update IDs and tokens as needed
- Check response status codes and data

## 🔍 Response Validation

Each request includes test scripts that validate:

- **Status Codes** - Correct HTTP status codes
- **Response Structure** - Expected JSON structure
- **Data Integrity** - Proper field values and types
- **Error Handling** - Appropriate error responses

## 📊 Expected Response Formats

### Success Response

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {
    // Response data here
  }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error description",
  "error": {
    "code": "ERROR_CODE",
    "message": "Error message",
    "details": "Additional error details"
  }
}
```

### Validation Error Response

```json
{
  "success": false,
  "message": "Validation failed",
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Request validation failed"
  },
  "data": {
    "validation_errors": [
      {
        "field": "title",
        "message": "Title is required"
      }
    ]
  }
}
```

## 🐛 Troubleshooting

### Common Issues

1. **Connection Refused**

   - Verify Posts API is running on correct port
   - Check `baseUrl` in environment configuration

2. **401 Unauthorized**

   - Verify JWT token is valid and not expired
   - Check Users API is accessible
   - Ensure token is properly formatted in Bearer auth

3. **404 Not Found**

   - Verify API endpoints are correct
   - Check if post IDs exist in the database
   - Ensure route patterns match (numeric IDs only)

4. **403 Forbidden**
   - Verify you're the author of the post you're trying to modify
   - Check JWT token contains correct user information

### Debug Steps

1. **Check API Status**

   ```
   GET {{baseUrl}}/health
   ```

2. **Verify Authentication**

   ```
   POST {{usersApiUrl}}/auth/me
   Authorization: Bearer your-jwt-token
   ```

3. **Check Database**
   - Verify database connection in API logs
   - Check if posts exist in the database

## 🔧 Customization

### Adding New Tests

1. Create new `.bru` files in appropriate folders
2. Follow existing patterns for structure
3. Include proper documentation and test scripts
4. Update sequence numbers for ordering

### Environment Configuration

- Add new environments in the `environments/` folder
- Update URLs and variables as needed
- Switch environments in Bruno interface

## 📋 API Endpoints Reference

| Method | Endpoint                          | Auth Required | Description               |
| ------ | --------------------------------- | ------------- | ------------------------- |
| GET    | `/health`                         | No            | Health check              |
| GET    | `/`                               | No            | API information           |
| GET    | `/api/v1/posts`                   | No            | Get all posts             |
| GET    | `/api/v1/posts/{id}`              | No            | Get post by ID            |
| GET    | `/api/v1/posts/author/{authorId}` | No            | Get posts by author       |
| POST   | `/api/v1/posts`                   | Yes           | Create new post           |
| PUT    | `/api/v1/posts/{id}`              | Yes           | Update post (author only) |
| DELETE | `/api/v1/posts/{id}`              | Yes           | Delete post (author only) |

## 📚 Additional Resources

- [Posts API Documentation](../README.md)
- [Users API Repository](https://github.com/drirodri/users-api)
- [Bruno Documentation](https://docs.usebruno.com/)
- [API Architecture Guide](../GUIDE.md)
