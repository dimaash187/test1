build:
	go build -o main main.go

run:
	go run main.go

test: ## run tests
	go test ./handlers -v