.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o email-service cmd/server/*
	docker build -t shippy-service-email -f ./build/package/Dockerfile.server .

.PHONY: run
run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-email
