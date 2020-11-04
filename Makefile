run:
	go run main.go
test:
	go test -cover ./...

cover:
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o coverage.html

swaggen:
	swag init