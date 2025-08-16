FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./main.go

FROM scratch

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]