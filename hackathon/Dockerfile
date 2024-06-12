FROM golang:1.22.2 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -v main.go
EXPOSE 8080
CMD ["./main"]