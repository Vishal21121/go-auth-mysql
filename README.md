# Go Auth MySQL

## Overview

Go Auth MySQL is a simple, yet powerful, authentication system built with Go and MySQL. It provides a robust foundation for user authentication, including registration, login, and password management. This project is designed to be easy to integrate into any Go-based web application, offering a secure and efficient way to manage user access.

## Features

- **User Registration**: Allows new users to sign up with their email and password.
- **User Login**: Authenticates users based on their email and password.
- **Password Management**: Supports password hashing and verification to ensure security.

## Prerequisites

- Go (version 1.22 or later)
- MySQL (version 8.0.30 or later)

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/Vishal21121/go-auth-mysql.git
   cd go-auth-mysql
   ```

2. **Setup MySQL Database**

   Create a new MySQL database and user for the application. Update the `.env` file with your database credentials.

   ```env
   USER=
   PASSWORD=
   HOST=
   PORT=
   DB_NAME=
   ```

3. **Run application**

   `go run main.go`

## Usage

The application exposes several RESTful API endpoints for user authentication and management. You can use tools like Postman or curl to interact with these endpoints.

- **Register a New User**

```http
POST /api/v1/users/register
Content-Type: application/json

{
   "name":"sudouser",
   "email": "user@example.com",
   "password": "your_password"
}
```

- **Login a User**

```http
POST /api/v1/users/login
Content-Type: application/json

{
   "email": "user@example.com",
   "password": "your_password"
}
```
