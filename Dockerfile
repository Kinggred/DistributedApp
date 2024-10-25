FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy the entire project
COPY . .
RUN go build -o tic-tac-toe ./main.go
CMD ["/app/tic-tac-toe"]

