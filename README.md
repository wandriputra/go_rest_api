# go_rest_api

This is a boilerplate for building RESTful APIs using Go and Fiber. It includes a basic setup for connecting to a PostgreSQL database and demonstrates how to structure your project for scalability and maintainability.

## Features

- Fiber for fast and flexible web framework
- GORM for ORM (Object-Relational Mapping)
- PostgreSQL as the database
- Docker and Docker Compose for containerization
- Air for hot reload during development
- Basic user service example

## Project Structure
go_rest_api/ 
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
├── go.sum
├── main.go
├── README.md 
└── src/ 
    ├── domains/ 
    │     └── user.go 
    ├── services/ 
    │     └── user.go 
└── public/ 
    └── input.json


## Getting Started

### Prerequisites

- Go 1.17 or later
- Docker
- Docker Compose

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/go_rest_api.git
   cd go_rest_api
2. Create a .env file in the root directory with the following content:
```
DB_HOST=db
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
```
3. Build and run the application using Docker Compose:
```bash
docker-compose up --build
```

Using Air for Hot Reload
Air is a live reload tool for Go applications. It watches for changes in your source code and automatically reloads the application.
Install Air:
curl -fLo air https://raw.githubusercontent.com/cosmtrek/air/master/bin/install.sh
chmod +x air
mv air /usr/local/bin

Run the application with Air:

air

This will start the application and watch for any changes in your source code. When changes are detected, Air will automatically reload the application.

Usage
The boilerplate includes a basic user service example. You can extend it by adding more services and routes as needed.

Contributing
Feel free to submit issues and pull requests for improvements and new features.

License
This project is licensed under the MIT License.


