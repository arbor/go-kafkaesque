all: run coverage test

.PHONY: test
test:
	# Ignore doc_test.go because example would make a HTTP request
	# and fail.
	unset GO111MODULE && go test -race -v ./... -tags doc

.PHONY: coverage
coverage: test
	$(eval GO111MODULE := off)
	go get github.com/axw/gocov/gocov
	gocov test -tags doc | gocov report
	$(eval GO111MODULE := on)
	go mod tidy
	go mod vendor

.PHONY: run
run:
	go fmt ./...
