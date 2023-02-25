build:
	env GOOS=linux GOARCH=amd64 go build -o bin/publicip-amd64 ./main.go
