build:
	@go fmt main.go
	@go build -o bin/hotel.exe

run: build
	@./bin/hotel

test:
	@go test -v ./..
