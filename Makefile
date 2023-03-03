BUILD_PATH=./cmd/balance-api


all: clean build run 

.PHONY: run
run:
	go run $(BUILD_PATH)/main.go


.PHONY: build
build:
	go build -o  $(BUILD_PATH) $(BUILD_PATH)/


.PHONY: docker-compose
docker-compose:
	docker-compose up 

.PHONY: docker-down
docker-down:
	docker-compose down 

.PHONY: docker-build
docker-build:
	docker build -t balance-api  .

.PHONY: docker-run 
docker-run:
	docker run -d -p 8080:8080 --name balanceAPI balance-api

.PHONY: docker-stop 
docker-stop: 
	docker stop balance-api

.PHONY: unit-test
unit-test:
	go test ./...


.PHONY: integrations-test
integrations-test:


.PHONY: clean 
clean:
