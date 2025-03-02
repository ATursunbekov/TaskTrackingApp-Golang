TaskTrackingApp-Golang

**TaskTrackingApp-Golang** is a task management system built with Go (Golang), designed to help users efficiently track and manage their tasks. This application provides functionalities such as user authentication, task creation, updating, deletion, and retrieval.

## Features

- **User Authentication**: Secure user registration and login using JWT (JSON Web Tokens).
- **Task Management**: Create, read, update, and delete tasks.
- **RESTful API**: Provides a RESTful API for interaction with the frontend or other services.
- **Swagger Documentation**: Integrated API documentation using Swagger.

### Prerequisites

- **Go**: Ensure you have Go installed. You can download it from the [official website](https://golang.org/dl/).
- **PostgreSQL**: The application uses PostgreSQL as its database. Install it from [here](https://www.postgresql.org/download/).

## Project Structure

```
TaskTrackingApp-Golang/
├── cmd/
│   ├── main.go          # Entry point of the application
│   └── migrate.go       # Database migration script
├── configs/             # Configuration files
├── docs/                # Swagger documentation files
├── pkg/
│   ├── handler/         # HTTP handlers
│   ├── repository/      # Database interactions
│   └── service/         # Business logic
├── schema/              # Database schema and migrations
├── .env                 # Environment variables
├── go.mod               # Go module file
└── README.md            # Project documentation
```
