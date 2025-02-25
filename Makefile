run-server:
	go run ./cmd/server/main.go

run-db:
	docker-compose up -d

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migration-up:
	@goose up

migration-down:
	@goose down
