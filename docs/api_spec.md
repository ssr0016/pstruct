# API Specification

## Task Management API

### Create Task
- **Endpoint**: POST /tasks
- **Request Body**: `{"title": "string", "description": "string", "status": "string"}`

### Get Task By ID
- **Endpoint**: GET /tasks/:id

### Update Task
- **Endpoint**: PUT /tasks/:id
- **Request Body**: `{"title": "string", "description": "string", "status": "string"}`

### Delete Task
- **Endpoint**: DELETE /tasks/:id

### Get All Tasks
- **Endpoint**: GET /tasks

## User Management API

### Create User
- **Endpoint**: POST /users
- **Request Body**: `{"username": "string", "password": "string"}`

### Get User By ID
- **Endpoint**: GET /users/:id

### Update User
- **Endpoint**: PUT /users/:id
- **Request Body**: `{"username": "string", "password": "string"}`

### Delete User
- **Endpoint**: DELETE /users/:id

### Get All Users
- **Endpoint**: GET /users
