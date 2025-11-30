
## test: run all tests that use example input
.PHONY: test-examples
test-examples:
	go test -v -race -test.short ./...

## test: run all tests
.PHONY: test-all
test-all:
	go test -v -race ./...