test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

swag:
	go install github.com/swaggo/swag/cmd/swag@latest && swag init  --output swagger/
