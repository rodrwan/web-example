.PHONY: run build

PORT=9091

run:
	@echo "Running server..."
	go generate ./... && PORT=$(PORT) go run main.go

build:
	@echo "Building server..."
	go generate ./... && go build -o bin/server main.go

clean:
	@echo "Cleaning up..."
	rm -rf bin/*