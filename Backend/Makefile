build:
	@go build -o bin/backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/backend

# migration:
# 	@migrate -path cmd/migrate/migrations -database "$(MONGODB_URI)" up

# migrate-up:
# 	@go run cmd/migrate/main.go up

# migrate-down:
# 	@go run cmd/migrate/main.go down