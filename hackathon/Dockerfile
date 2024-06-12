FROM golang:1.22.2 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -v *.go
EXPOSE 8080
CMD ["go", "run", "main.go"]