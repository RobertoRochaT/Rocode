# RoCode Backend

Backend API for RoCode - A LeetCode pattern-based learning platform.

## Tech Stack

- **Go 1.21+**
- **PostgreSQL**
- **Gin** - Web framework
- **JWT** - Authentication

## Features

- User authentication (register/login)
- Progress tracking for LeetCode problems
- 94 programming patterns organized by category
- RESTful API endpoints
- CORS support for frontend integration

## Setup

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12+

### Environment Variables

Create a `.env` file or set the following environment variables:

```bash
DATABASE_URL=postgres://postgres:postgres@localhost:5432/rocode?sslmode=disable
JWT_SECRET=your-secret-key-change-in-production
PORT=8080
```

### Database Setup

The application will automatically create the required tables on startup:
- `users` - User accounts
- `problem_progress` - User progress tracking

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run .
```

The server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

#### Auth
- `POST /api/register` - Register a new user
- `POST /api/login` - Login and receive JWT token
- `GET /api/patterns` - Get all programming patterns
- `GET /api/problems` - Get all problems (optionally filtered by pattern_id)

### Protected Endpoints (Require JWT Token)

Include `Authorization: Bearer <token>` header in requests.

#### Progress
- `GET /api/progress` - Get user's progress
- `POST /api/progress` - Update problem progress
- `GET /api/stats` - Get user statistics

## API Examples

### Register
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "johndoe",
    "password": "password123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### Update Progress
```bash
curl -X POST http://localhost:8080/api/progress \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "problem_number": "1",
    "pattern_id": 11,
    "completed": true,
    "difficulty": "Easy",
    "notes": "Solved using hash map"
  }'
```

## Docker

Build and run with Docker:

```bash
docker build -t rocode-backend .
docker run -p 8080:8080 -e DATABASE_URL=<your-db-url> rocode-backend
```

## Project Structure

```
backend/
├── main.go           # Application entry point & database setup
├── auth.go           # Authentication handlers & middleware
├── handlers.go       # API handlers for patterns, problems, progress
├── problems.go       # Problem data (all LeetCode problems)
├── go.mod            # Go module dependencies
├── go.sum            # Go module checksums
├── Dockerfile        # Docker configuration
└── README.md         # This file
```

## Database Schema

### users
- `id` - Primary key
- `email` - Unique user email
- `username` - Unique username
- `password_hash` - Bcrypt hashed password
- `created_at` - Account creation timestamp

### problem_progress
- `id` - Primary key
- `user_id` - Foreign key to users
- `problem_number` - LeetCode problem number
- `pattern_id` - Pattern ID
- `completed` - Boolean completion status
- `difficulty` - Problem difficulty (Easy/Medium/Hard)
- `notes` - User notes
- `completed_at` - Completion timestamp
- `created_at` - Record creation timestamp
- `updated_at` - Last update timestamp

## License

MIT