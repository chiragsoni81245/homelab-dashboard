build:
	@go build -o bin/homelab-dashboard ./main.go

run: build
	@./bin/homelab-dashboard start $(ARG)

test:
	@go test ./... -v --race
