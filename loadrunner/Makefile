.PHONY: help
help:
	@echo "Available make commands:"
	@cat Makefile | grep '^[a-z][^:]*:' | cut -d: -f1 | sort | sed 's/^/  /'

rundep=go run -modfile ../misc/devdeps/go.mod

.PHONY: build
build: build.loadrunner

build.loadrunner:;    go build -o build/loadrunner   ./cmd

.PHONY: install
install: install.loadrunner

install.loadrunner:;    go install ./cmd

PHONY: clean
clean:
	rm -rf build

.PHONY: lint
lint:
	$(rundep) github.com/golangci/golangci-lint/cmd/golangci-lint run --config ../.github/golangci.yml ./...

.PHONY: fmt
GOFMT_FLAGS ?= -w
fmt:
	$(rundep) mvdan.cc/gofumpt $(GOFMT_FLAGS) .

########################################
# Test suite
.PHONY: test
test: _test.loadrunner

GOTEST_FLAGS ?= -v -p 1 -timeout=30m

_test.loadrunner:;      go test $(GOTEST_FLAGS) ./cmd

