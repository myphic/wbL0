FROM golang:1.18
WORKDIR /wildberriesL0
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]