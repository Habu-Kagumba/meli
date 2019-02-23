.PHONY: build

default: build

build:
	@echo ">>>> Building Docker <<<<"
	docker build -t consignment-cli .

run:
	@echo ">>>> Running Application <<<<"
	docker run -e MICRO_REGISTRY=mdns consignment-cli
