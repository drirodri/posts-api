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
│   │   └── user.go          # Reference model for Users API
│   ├── repository/
│   │   └── post_repository.go
│   ├── services/
│   │   ├── post_service.go
│   │   └── user_service.go  # HTTP client to Users API
│   ├── handlers/
│   │   └── post_handler.go
│   ├── middleware/
│   │   ├── auth.go         # JWT validation
│   │   └── cors.go
│   └── dto/
│       ├── create_post.go
│       ├── update_post.go
│       └── post_response.go
├── pkg/
│   └── utils/
│       └── response.go      # Standardized responses
├── go.mod
├── go.sum
├── .env
└── README.md