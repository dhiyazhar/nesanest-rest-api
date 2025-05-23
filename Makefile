BINARY_NAME=nesanest-api
BUILD_PATH=cmd/main.go
MIGRATION_PATH=./migrations
DB_URL=postgresql://azhar:azhar@localhost:5432/nesanest?sslmode=disable

.PHONY: build migrate run clean

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(BUILD_PATH)
	@echo "Build complete."

migrate:
	@echo "Running database migrations..."
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up
	@echo "Migrations complete."

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "Cleanup complete."
