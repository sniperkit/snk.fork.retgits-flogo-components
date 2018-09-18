.PHONY: deps prepare build install

deps:
	@go get github.com/retgits/ghhelper
	@go get github.com/retgits/gc
	@go get github.com/retgits/fdio
	@go get github.com/retgits/flogowebhelper
	@go get github.com/retgits/fgen
	@go get github.com/TIBCOSoftware/flogo-cli/...

prepare: deps
	@amalgomate --config config.yml --output-dir fdev --pkg main

build:
	@go build -o ./bin/fdev ./fdev/main.go

install:
	@go install github.com/sniperkit/snk.fork.retgits-flogo-components/command/repack/fdev
