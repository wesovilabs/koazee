# Go Tools
GO  = GO111MODULE=on go
all: fmt check build test info
clean:
	rm -f coverage.txt
	rm -rr docs/_site
	rm -r docs/.sass-cache
	rm -rf vendor
deps:
	${GO} mod vendor
	${GO} mod download
test:
	${GO} test  -v ./...
test-coverage:
	${GO} test -race -coverprofile=coverage.txt -covermode=atomic ./...
fmt:
	GO111MODULE=on ${GO} fmt .
	goimports -w .
check: fmt
	gometalinter --deadline 60s --enable lll --enable goimports --disable errcheck --disable gotype --disable golint --disable aligncheck
check-fast:
	gometalinter --enable goimports --enable lll --disable errcheck --disable gotype  --disable golint --fast
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
	${GO} get github.com/divan/depscheck
	GO111MODULE=off go get github.com/alecthomas/gometalinter
	GO111MODULE=off gometalinter --install --force
site:
	cd docs; \
	bundle install; \
	jekyll serve
