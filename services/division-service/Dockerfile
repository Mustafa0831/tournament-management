FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o division-service .

EXPOSE 8082

CMD ["./division-service"]
