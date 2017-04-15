SOURCES := main.go $(WILDCARD cmds/*.go)

go-otp: $(SOURCES)
	go build -o go-otp .

gofmt:
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")
	gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")
