APP_NAME = "json-to-go"

run: build
	@./bin/$(APP_NAME) --input input.json

build:
	@go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)
