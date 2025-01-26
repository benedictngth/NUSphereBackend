# NUSphereBackend

## Overview

Backend repository for NUSphere - a social networking platform for NUS students.

## Tech Stack

- Golang with Go-Gin framework
- pgx driver for PostgreSQL
- Postgreql
- JWT Authentication with Golang-jwt
- HTTP-Only Cookies for secure token storage

## Setup

1. Clone the repository

```bash
git clone https://github.com/your-username/NUSphereBackend.git
git checkout localDevelopment
```

2. Install go dependencies

```bash
go mod download
```

3. Set up environment variables
   Create a `.env` file with:

```sh
DATABASE_URL=postgresql://<username>:<password>@localhost:5432/<dbname>?sslmode=disable
 POSTGRES_USER=<username>
 POSTGRES_PASSWORD=<password>
 POSTGRES_DB=<dbname>
 POSTGRES_HOST=db
 POSTGRES_PORT=5432
 POSTGRES_SSLMODE=disable
 JWT_SECRET=<your-jwt-secret>
 PORT=8080
```

4. Start the server

```bash
go run app.go
```

## Database Schema

### ID System

- Each entity uses two ID types:
  - `id`: Internal auto-incrementing primary key
  - `public_id`: External NanoID for public API exposure

### Architecture Pattern

Each database table follows a three-layer architecture:

1. **Repository Layer**

   - Handles direct database operations
   - Implements CRUD operations
   - Uses pgx for database interactions

2. **Service Layer**

   - Contains business logic
   - Validates data
   - Coordinates between repository and router

3. **Router Layer**
   - Manages HTTP endpoints
   - Handles request/response
   - Implements middleware
   - Uses Gin framework

Example structure for `users`:

```
├── users/
│   ├── userRepository.go
│   ├── authService.go
│   ├── router.go
│   ├── middlewares.go
│   ├── constants.go
│   └── models.go
```

## API Documentation

### Users

- POST /api/users/register - Register new user
- POST /api/users/login - User login
- GET /api/users/auth - Check authentication status
- GET /api/users - Get all users
- GET /api/users/authUser - Get authenticated user details
- POST /api/users/logout - Logout user

### Posts

- POST /api/posts/create - Create new post
- GET /api/posts - Get all posts
- GET /api/posts/:id - Get post by ID
- PUT /api/posts/edit/:id - Edit post
- DELETE /api/posts/delete/:id - Delete post

### Comments

- POST /api/comments/create - Create new comment
- GET /api/comments - Get comments by post ID
- GET /api/comments/:id - Get comment by ID
- PUT /api/comments/edit/:id - Edit comment
- DELETE /api/comments/delete/:id - Delete comment

### Categories

- POST /api/categories/create - Create new category
- GET /api/categories - Get all categories
