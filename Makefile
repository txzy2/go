APP_NAME = first_go
TAG = latest
BINARY_NAME = tmp/main
CMD_PATH = ./cmd/first

test:
	go test -v ./...

build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

run: build
	./$(BINARY_NAME)

docker-build:
	docker build -t $(APP_NAME):$(TAG) .

docker-run:
	docker run --rm $(APP_NAME):$(TAG)

# Комбинация build + run
br: build run

# Комбинация docker-build + docker-run
dr: docker-build docker-run

.PHONY: test build run docker-build docker-run br dr
