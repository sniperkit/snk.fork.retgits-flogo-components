.PHONY: deps prepare build install

deps:
	@go get github.com/genuinetools/apk-file
	@go get github.com/opennota/check/cmd/varcheck
	@go get github.com/mdempsky/unconvert
	@go get github.com/gordonklaus/ineffassign
	@go get github.com/nmiyake/govet
	@go get github.com/golang/lint/golint
	@go get github.com/kisielk/errcheck
	@go get github.com/palantir/distgo
	@go get github.com/palantir/bulldozer
	@go get github.com/retgits/ghhelper
	@go get github.com/retgits/gc
	@go get github.com/jesseduffield/lazygit
	@go get github.com/retgits/fdio
	@go get github.com/retgits/flogowebhelper
	@go get github.com/retgits/fgen
	@go get github.com/TIBCOSoftware/flogo-cli/...

prepare:
	@amalgomate --config config.yml --output-dir f++ --pkg main

build:
	@go build -o ./bin/f++ ./f++/main.go

install:
	@go install github.com/sniperkit/snk.fork.retgits-flogo-components/command/repack/f++
