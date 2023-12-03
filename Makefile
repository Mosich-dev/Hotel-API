build:
	go fmt main.go
	go build -o hotel.exe

run: build
	./hotel
