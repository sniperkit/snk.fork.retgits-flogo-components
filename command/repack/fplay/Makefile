.PHONY: deps prepare build install

deps:
	@go get github.com/TIBCOSoftware/flogo-cli/...
	@go get github.com/retgits/flogowebhelper
	@go get github.com/huytd/playgo

prepare: deps
	@amalgomate --config config.yml --output-dir fplay --pkg main

build:
	@go build -o ./bin/fplay ./fplay/main.go

install:
	@go install github.com/sniperkit/snk.fork.retgits-flogo-components/command/repack/fplay
