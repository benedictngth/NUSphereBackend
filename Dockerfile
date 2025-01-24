#purely for go application
# Stage 1: Build the Go application

# syntax=docker/dockerfile:1.3
FROM golang:1.23 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Copy .env file
# COPY .env .env

RUN --mount=type=secret,id=env \
    # 1) Make the secret file available
    # 2) Export the variables so they're available in this RUN step
    #    cat /run/secrets/env | xargs is a quick hack to turn KEY=val lines into env variables.
    # 3) Do something with them (like install packages, private repo clone, etc.)
    sh -c "export $(cat /run/secrets/env | xargs) && \
    echo \"The DATABASE_URL is: $DATABASE_URL\" && \
    echo \"The JWT KEY starts with: $JWT_SECRET\""


RUN go build -o app

# Stage 2: Run the Go application
# RUN apk --no-cache add ca-certificates
COPY migrate.sh .
RUN chmod +x migrate.sh
EXPOSE 8080
CMD ["./app"]
