.PHONY: build
build:
	protoc -I . --go_out=. \
		--micro_out=. \
		proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o vessel-service cmd/server/main.go
	docker build -t shippy-service-vessel -f ./build/package/Dockerfile.server .

.PHONY: run
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-vessel
