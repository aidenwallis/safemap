# Run unit tests:
test:
	@go test -v

coverage:
	@go test -coverprofile=coverage.out -covermode=count
	@go tool cover -html=coverage.out -o coverage.html
