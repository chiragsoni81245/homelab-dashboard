FROM golang:latest AS builder

WORKDIR /app

# Install Required Go Packages
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1 GOOS=linux
RUN go build \
    -tags "sqlite_omit_load_extension" \
    -ldflags='-linkmode external -extldflags "-static"' \
    -o /bin/homelab-dashboard

FROM scratch AS runner

COPY --from=builder /bin/homelab-dashboard /bin/homelab-dashboard

EXPOSE 8080

CMD ["/bin/homelab-dashboard", "start"]
