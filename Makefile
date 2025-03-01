lint:
	@go vet ./...

build-gin:
	@echo "Building..."
	@go build -o bin/gin cmd/gin/main.go

watch-gin:
	@air -c gin.air.toml

pre-commit:
	@pre-commit autoupdate && pre-commit install
