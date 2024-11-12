FROM golang:1.23 as builder

RUN apt-get update && apt-get install -y \
    pkg-config \
    libgl1-mesa-dev \
    xorg-dev \
    libglu1-mesa-dev

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy the entire project
COPY . .
RUN go build -o tic-tac-toe ./main.go
CMD ["/app/tic-tac-toe", "-mode=server_only"]

