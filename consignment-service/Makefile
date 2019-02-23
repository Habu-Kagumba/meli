.PHONY: grpc-build docker-build

default: build

build: grpc-build docker-build

grpc-build:
	@echo ">>>> Building gRPC <<<<"
	protoc --proto_path=$(GOPATH)/src/github.com/Habu-Kagumba/meli:. --micro_out=. --go_out=. proto/consignment/consignment.proto

docker-build:
	@echo ">>>> Building Docker <<<<"
	docker build -t consignment-service .

run:
	@echo ">>>> Running Application <<<<"
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service
