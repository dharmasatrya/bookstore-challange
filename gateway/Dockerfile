FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]