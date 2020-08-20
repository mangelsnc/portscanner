build:
	go build -o dist/portscanner main.go

run:
	go run main.go connectScan 127.0.0.1
