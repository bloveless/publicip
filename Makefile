IMAGE_TAG=0.1.3

run:
	go run ./main.go

build:
	build -o bin/publicip ./main.go

push:
	docker buildx build --push --platform linux/amd64,linux/arm64 -t bloveless/publicip:$(IMAGE_TAG) .

