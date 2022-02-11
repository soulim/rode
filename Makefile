BIN_DIR = ./bin
CMD_DIR = ./cmd/rode

VERSION = 0.1.0
LD_FLAGS = -ldflags=-X=main.version=$(VERSION)

GO = go

$(BIN_DIR):
	mkdir -p $@

$(BIN_DIR)/rode: check
	$(GO) build $(LD_FLAGS) -o $@ $(CMD_DIR)

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: test
test:
	$(GO) test ./...

.PHONY: check
check: | fmt vet test
