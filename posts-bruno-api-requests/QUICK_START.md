# Quick Start Guide - Posts API Testing

## 🚀 5-Minute Setup

### Step 1: Install Bruno

Download Bruno from [usebruno.com](https://www.usebruno.com/) and install it.

### Step 2: Open Collection

1. Open Bruno
2. Click "Open Collection"
3. Navigate to this folder: `posts-bruno-api-requests`
4. Click "Open"

### Step 3: Start APIs

```bash
# Terminal 1: Start Posts API
cd c:\Users\adria\Desktop\Programacao\Go\posts-api
go run cmd/server/main.go

# Terminal 2: Start Users API (for authentication)
# Follow Users API setup instructions
```

### Step 4: Get JWT Token

1. Register/Login to Users API to get a JWT token
2. Replace `your-jwt-token-here` in protected requests

### Step 5: Run Tests

1. Select "Local" environment
2. Start with "Health Check"
3. Run "Get All Posts"
4. Test "Create Post" (requires valid token)

## 🧪 Essential Test Sequence

### Basic Flow (No Auth)

1. ✅ Health Check
2. ✅ Root Endpoint
3. ✅ Get All Posts
4. ✅ Get Post by ID (if posts exist)

### CRUD Flow (Requires Auth)

1. 🔐 Create Post
2. 📖 Get Created Post
3. ✏️ Update Post
4. 🗑️ Delete Post

### Error Testing

1. ❌ Invalid Token
2. ❌ Missing Token
3. ❌ Validation Errors
4. ❌ Invalid IDs

## 📝 Common Test Data

### Create Post Body

```json
{
  "title": "Test Post Title",
  "content": "This is test content for the post."
}
```

### Update Post Body

```json
{
  "title": "Updated Title",
  "content": "Updated content with more information."
}
```

## 🔧 Troubleshooting

| Issue              | Solution                                    |
| ------------------ | ------------------------------------------- |
| Connection refused | Check if Posts API is running on port 3334  |
| 401 Unauthorized   | Update JWT token in protected requests      |
| 404 Not Found      | Verify post ID exists or check endpoint URL |
| 403 Forbidden      | Ensure you're the author of the post        |

## 📞 Need Help?

- Check the full [README.md](README.md) for detailed instructions
- Review API logs for error details
- Verify environment configuration
- Test APIs individually before running workflows
