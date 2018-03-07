build:
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/pkg/errors
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go