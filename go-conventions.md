# Go Variable Naming Conventions

## Common Abbreviations and Their Meanings

### Configuration & Setup

- `cfg` - Configuration
- `config` - Configuration (full word)
- `env` - Environment
- `dsn` - Data Source Name (database connection string)
- `conn` - Connection
- `db` - Database
- `srv` - Server
- `app` - Application

### Error Handling

- `err` - Error
- `ok` - Boolean flag (typically from map lookups or type assertions)

### HTTP & Web

- `req` - Request
- `res` - Response
- `resp` - Response (alternative)
- `w` - http.ResponseWriter
- `r` - \*http.Request
- `mux` - HTTP request multiplexer/router
- `handler` - HTTP handler function
- `middleware` - HTTP middleware
- `router` - HTTP router

### Data Transfer & Models

- `dto` - Data Transfer Object
- `model` - Data model
- `entity` - Database entity
- `repo` - Repository
- `svc` - Service
- `dao` - Data Access Object

### Context & Concurrency

- `ctx` - Context (context.Context)
- `cancel` - Context cancellation function
- `wg` - Wait group (sync.WaitGroup)
- `ch` - Channel
- `mu` - Mutex (sync.Mutex)
- `rwmu` - Read-write mutex (sync.RWMutex)

### Common Data Types

- `str` - String
- `num` - Number
- `cnt` - Count
- `idx` - Index
- `len` - Length
- `cap` - Capacity
- `ptr` - Pointer
- `val` - Value
- `key` - Key (in key-value pairs)

### Files & I/O

- `f` - File
- `fd` - File descriptor
- `buf` - Buffer
- `data` - Raw data
- `b` - Byte slice ([]byte)
- `n` - Number of bytes read/written

### Time & Scheduling

- `t` - Time
- `dur` - Duration
- `timeout` - Timeout duration
- `deadline` - Deadline time
- `ticker` - Time ticker
- `timer` - Timer

### Collections & Iteration

- `items` - Collection of items
- `list` - List/slice
- `arr` - Array
- `m` - Map
- `s` - Set or slice
- `i`, `j`, `k` - Loop indices
- `v` - Value in range loops
- `k` - Key in range loops

### Authentication & Security

- `auth` - Authentication
- `token` - Authentication token
- `jwt` - JSON Web Token
- `hash` - Hash value
- `salt` - Password salt
- `creds` - Credentials
- `user` - User object
- `role` - User role
- `perm` - Permission

### Logging & Monitoring

- `log` - Logger
- `msg` - Message
- `lvl` - Log level
- `trace` - Trace information
- `metric` - Metric value

### Business Logic

- `svc` - Service
- `mgr` - Manager
- `ctrl` - Controller
- `proc` - Processor
- `validator` - Validator
- `parser` - Parser
- `encoder` - Encoder
- `decoder` - Decoder

### Database & Queries

- `tx` - Transaction
- `stmt` - Prepared statement
- `rows` - Query result rows
- `row` - Single database row
- `query` - SQL query string
- `args` - Query arguments
- `result` - Query result

### Configuration Examples from Your Code

```go
cfg := &AppConfig{...}           // Configuration object
dsn := fmt.Sprintf("host=%s...", // Data Source Name for database
err := gorm.Open(...)            // Error from operation
DB, err := gorm.Open(...)        // Database connection and error
```

## Naming Best Practices

### Do:

- Use short, clear abbreviations for common concepts
- Be consistent across your codebase
- Use descriptive names for important business logic
- Follow Go conventions (camelCase for private, PascalCase for public)

### Don't:

- Use single letters except for short-lived variables (i, j, k in loops)
- Create your own abbreviations that others won't understand
- Use overly long variable names for simple operations
- Mix naming conventions within the same project

## Common Patterns

### Error Handling Pattern

```go
result, err := someFunction()
if err != nil {
    return err
}
```

### HTTP Handler Pattern

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // w = response writer, r = request
}
```

### Context Pattern

```go
func doSomething(ctx context.Context) error {
    // ctx = context for cancellation/timeouts
}
```

### Range Loop Pattern

```go
for k, v := range items {
    // k = key, v = value
}

for i, item := range items {
    // i = index, item = value
}
```
