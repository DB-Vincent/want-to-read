# Want to Read

A web application to manage your reading list and track books you want to read.

## Project Structure

```
want-to-read/
├── backend/                 # Golang backend
│   ├── cmd/                # Application entry points
│   │   └── api/           # API server
│   ├── internal/          # Private application code
│   │   ├── models/       # Data models
│   │   ├── handlers/     # HTTP handlers
│   │   ├── services/     # Business logic
│   │   └── database/     # Database operations
│   ├── pkg/              # Public library code
│   ├── api/              # API definitions
│   └── configs/          # Configuration files
└── frontend/              # Angular frontend
    ├── src/              # Source code
    ├── angular.json      # Angular configuration
    └── package.json      # Frontend dependencies
```

## Prerequisites

- Go 1.21 or later
- Node.js and npm
- Angular CLI
- PostgreSQL

## Setup Instructions

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the backend directory with your database configuration:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=want_to_read
   ```

4. Run the server:
   ```bash
   go run cmd/api/main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   ng serve
   ```

The application will be available at:
- Frontend: http://localhost:4200
- Backend API: http://localhost:8080

## Features

- Add books to your reading list
- Track book details (title, author, description)
- Set reading priorities
- Mark books as read
- Search and filter your reading list

## API Endpoints

- `GET /health` - Health check endpoint
- More endpoints coming soon...

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
