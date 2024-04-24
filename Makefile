APP_NAME = "json-to-go"

run: build
	@echo "Running..."
	@./bin/$(APP_NAME)

build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)
