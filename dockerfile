# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN rm -rf ./build
RUN go build -o ./build/app ./cmd/app/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:3.14

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/build .

# Expose the port that the application will run on
EXPOSE 8080

CMD ["./app"]