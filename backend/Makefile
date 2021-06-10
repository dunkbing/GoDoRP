# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
BINARY_NAME=bin/backend
BINARY_UNIX=bin_unix/backend

all: test build
build:
	$(GO_BUILD) -o $(BINARY_NAME) -v
test:
	$(GO_TEST) -v ./...
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GO_BUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)
deps:
	$(GO_GET)

swagger:
	swag init
	@echo done
