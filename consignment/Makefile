.PHONY: server
server:
	protoc -I . --go_out=. \
		--micro_out=. \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o consignment-service cmd/server/main.go
	docker build -t shippy-service-consignment -f ./build/package/Dockerfile.server .

.PHONY: client
client:
	protoc -I . --go_out=. \
		--micro_out=. \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o consignment-cli cmd/cli/main.go
	docker build -t shippy-cli-consignment -f ./build/package/Dockerfile.cli .

.PHONY: run-server
run-server:
	docker run --rm -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 \
		shippy-service-consignment

.PHONY: run-client
run-client:
	docker run --rm shippy-cli-consignment
