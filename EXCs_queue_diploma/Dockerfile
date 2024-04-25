FROM golang:latest

WORKDIR /app

COPY . /app

COPY go.mod go.sum ./

RUN go mod download

# Build the Go app
RUN go build -o main ./cmd


EXPOSE 3000
ENTRYPOINT ["./main"]