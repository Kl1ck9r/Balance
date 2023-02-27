BUILD_PATH=./cmd/balance-api


.PHONY: run
run:
	go run $(BUILD_PATH)/main.go


.PHONY: build
build:
	go build -o  $(BUILD_PATH) $(BUILD_PATH)/


.PHONY: docker-compose
docker-compose:

.PHONY: docker-down
docker-down:


.PHONY: docker-build
docker-build:


.PHONY: docker-run 
docker-run:

.PHONY: docker-stop 
docker-stop: 

.PHONY: test
test:



.PHONY: clean 
clean:
