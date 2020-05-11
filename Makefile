export GO111MODULE=on

CMDS=$(filter-out internal, $(notdir $(wildcard cmd/*)))

### Local dev ----
.PHONY: test
test:
	go test --cover --race -v ./...

.PHONY: build
build: $(CMDS)

$(CMDS):
	go build  -o ./bin/$@ cmd/$@/*.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run ./...
