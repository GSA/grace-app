GOBIN := $(GOPATH)/bin
GOLANGCILINT := $(GOBIN)/golangci-lint

.PHONY: lint dependencies test
default: lint

test: lint
	go test -v -cover ./...

lint: dependencies
	$(GOLANGCILINT) run ./...

dependencies: $(GOLANGCILINT)

$(GOLANGCILINT):
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

go.sum: go.mod
	go mod tidy

go.mod:
	go mod init