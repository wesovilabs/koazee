# Go Tools
GO  = GO111MODULE=on go
all: fmt check build test info
clean:
	rm -f coverage.txt
	rm -rf vendor
deps:
	#${GO} mod vendor
	${GO} mod download
test:
	${GO} test  -v ./...
test-coverage:
	${GO} test -race -coverprofile=coverage.txt -covermode=atomic ./...
fmt:
	GO111MODULE=on ${GO} fmt ./...
check: fmt
	golangci-lint run
lint:
	golint
benchmark: fmt
	#${GO} run ./benchmark/main.go
	${GO} test -bench Benchmark.+ -run -Benchmark.+ -v ./benchmark/contains/...
info: fmt
	depscheck -totalonly -tests .
	golocc
std-info: fmt
	depscheck -stdlib -v .
install:
	${GO} get -u github.com/divan/depscheck
	${GO} install github.com/golangci/golangci-lint/cmd/golangci-lint
site:
	cd docs; \
	jekyll serve
doc:
	cd .hugo; \
	jekyll serve
