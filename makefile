build:
	go build -o bin/migration

run:
	go run main.go
	
producer:
	go run main.go producer