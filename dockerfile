FROM golang:1.25.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bloger ./cmd/main/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bloger .

COPY ./config/dev.yml ./config/dev.yml

EXPOSE 3000

# 运行应用
CMD ["./bloger"]
