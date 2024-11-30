# Use Go Alpine image
FROM golang:1.23.3-alpine

# Install necessary tools
RUN apk add --no-cache git bash

# Set up working directory
WORKDIR /app

# Clone the Air repository
RUN git clone https://github.com/cosmtrek/air.git /tmp/air

# Build the Air binary
WORKDIR /tmp/air
RUN go build -o /usr/bin/air .

# Return to the application directory
WORKDIR /app

# Copy dependency files and download modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the entire project
COPY . ./

# Expose the application port
EXPOSE 3000

# Use Air for hot reloading
CMD ["air"]