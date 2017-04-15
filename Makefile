SOURCES := main.go $(WILDCARD cmds/*.go)
VERSION := $(shell git describe --tags)

go-otp: $(SOURCES)
	go build -o go-otp .

deps:
	go get github.com/mitchellh/gox \
		github.com/tcnksm/ghr \
	code.google.com/p/rsc/qr \
	github.com/hgfischer/go-otp \
	github.com/maxmclau/gput \
	github.com/urfave/cli \
	github.com/xeodou/go-sqlcipher

gofmt:
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")
	gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

release:
	rm -rf dist/
	mkdir dist
	CGO_ENABLED=0 gox \
		-ldflags "-X main.version=$(VERSION) -X main.builddate=`date -u +%Y-%m-%dT%H:%M:%SZ`" \
		-output "dist/$(TARGET)_{{.OS}}_{{.Arch}}" \
		-os="linux"
	ghr -u erasche -replace $(VERSION) dist/
