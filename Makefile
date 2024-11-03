.PHONY: run

all:    
	@clear    
	@go fmt ./...    
	@go build -o bin/app ./cmd/main.go    
	@./bin/app

tidy:
	@clear && go mod tidy
	
run:    
	@clear    
	@go run ./cmd/main.go

clean:    
	@clear    
	@go clean    
	@rm -f bin/app

format:
	@clear    
	@go fmt ./...