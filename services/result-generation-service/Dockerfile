FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o result-generation-service .

EXPOSE 8085

CMD ["./result-generation-service"]
