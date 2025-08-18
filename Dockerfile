FROM golang:latest

WORKDIR /app

# Install Required Go Packages
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /homelab-dashboard

EXPOSE 8080

CMD ["/homelab-dashboard", "start"]
