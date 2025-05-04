FROM golang:1.24-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o backend .

# TODO:use environment variable instead of hardcode 
EXPOSE 8080


CMD ["./backend"]

