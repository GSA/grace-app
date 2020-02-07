GOBIN := $(GOPATH)/bin
GOLANGCILINT := $(GOBIN)/golangci-lint

.PHONY: lint dependencies test
default: lint

test: lint
	go test -v -cover ./...

lint:	precommit dependencies
	$(GOLANGCILINT) run ./...

dependencies: $(GOLANGCILINT)

$(GOLANGCILINT):
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

go.sum: go.mod
	go mod tidy

go.mod:
	go mod init

precommit:
ifeq ($(UNAME_S),Linux)
ifeq (,$(findstring Microsoft,$(uname -a)))
	dos2unix .github/hooks/pre-commit
endif
endif
ifneq ($(strip $(hooksPath)),.github/hooks)
	@git config --add core.hooksPath .github/hooks
endif
