.PHONY: lint

lint:
	gometalinter -I"linters\.go:" .

test: lint.go lint_test.go linters.go testdata/*
	go test .
