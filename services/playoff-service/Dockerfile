FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o playoff-service .

EXPOSE 8084

CMD ["./playoff-service"]
