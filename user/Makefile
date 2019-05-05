.PHONY: server
server:
	protoc -I . --go_out=. \
		--micro_out=. \
		proto/user/user.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user-service cmd/server/*
	docker build -t shippy-service-user -f ./build/package/Dockerfile.server .

.PHONY: client
client:
	protoc -I . --go_out=. \
		--micro_out=. \
		proto/user/user.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user-cli cmd/cli/main.go
	docker build -t shippy-cli-user -f ./build/package/Dockerfile.cli .

.PHONY: run-server
run-server:
	docker run --rm -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 \
		shippy-service-user

.PHONY: run-client
run-client:
	docker run --rm shippy-cli-user
