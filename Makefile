.PHONY: lint

lint:
	gometalinter --disable="gocyclo" -e"lint\.go" -e"lint_test\.go" .

test: lint.go lint_test.go linters.go testdata/*
	go test .
