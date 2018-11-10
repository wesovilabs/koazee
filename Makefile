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
	GO111MODULE=on ${GO} fmt .
	goimports -w .
check: fmt
	golangci-lint run
lint:
	golint
benchmark: fmt
	${GO} test -bench Benchmark.+ -run -Benchmark.+ -v
info: fmt
	depscheck -totalonly -tests .
	golocc
std-info: fmt
	depscheck -stdlib -v .
install:
	${GO} get -u github.com/divan/depscheck
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin v1.12.1
site:
	cd docs; \
	bundle install; \
	jekyll serve
