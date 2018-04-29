default: test

test:
	go test -cover -v $(shell go list ./... | grep -v vendor)

cover: depsdev
	goveralls -service=travis-ci

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

depsdev:
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go get github.com/golang/lint/golint
	go get github.com/motemen/gobump/cmd/gobump

.PHONY: default test deps cover
