#purely for go application
# Stage 1: Build the Go application
FROM golang:1.23 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Copy .env file
COPY .env .env
RUN go build -o app

# Stage 2: Run the Go application
# RUN apk --no-cache add ca-certificates
COPY migrate.sh .
RUN chmod +x migrate.sh
EXPOSE 8080
CMD ["./app"]
