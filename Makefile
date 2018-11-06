build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/ipsum ipsum/main.go
test:
	env GOOS=linux go clean --testcache && go test ./...
