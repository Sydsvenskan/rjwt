export VERSION ?= $(shell git describe --tags)

.PHONY: release
release:
	rm -fr build
	gox -os="linux darwin" -arch="amd64" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}-$(VERSION)"
	gzip build/*

.PHONY: get-tools
get-tools:
	go get -u -v github.com/mitchellh/gox
