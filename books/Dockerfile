FROM golang:1.23.2

WORKDIR /app

# Copy proto repo first
COPY proto-repo /app/proto-repo

# Copy users service files
COPY books /app/books

WORKDIR /app/books

# Update go.mod to use local path
RUN go mod edit -replace github.com/dharmasatrya/proto-repo=/app/proto-repo

# Download dependencies
RUN go mod download

# Build
RUN go build -o main .

EXPOSE 50052

CMD ["./main"]