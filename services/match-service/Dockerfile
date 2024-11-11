FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o match-service .

EXPOSE 8083

CMD ["./match-service"]
