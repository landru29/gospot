GOTEST             =gotest
ifeq (, $(shell which gotest))
    GOTEST=go test
endif


.PHONY: gospot gen

gospot:
	go build -o $@ ./cmd/...

clean:
	rm gospot

gen:
	GOFLAGS=-mod=mod go generate ./...

lint:
	golangci-lint run --timeout 10m0s --allow-parallel-runners $(params) ./...

lint-fix: param=--fix
lint-fix: lint

test:
	$(GOTEST) -coverprofile coverage.out ./... && go tool cover -func=coverage.out