FROM golang:1.23

WORKDIR /app

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build main.go

CMD ["./main"]