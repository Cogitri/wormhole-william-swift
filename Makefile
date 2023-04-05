.DEFAULT: build

build: Frameworks $(wildcard Package.*) $(wildcard Sources/*/*)
	swift build -c release

.PHONY: Frameworks
Frameworks: Frameworks/WormholeWilliamGo.xcframework

Frameworks/%.xcframework: Makefile tools $(wildcard Sources/*/*.go) $(shell which gomobile)
	mkdir -p Frameworks
	gomobile init
	gomobile bind -target=macos -x -o $@ ./Sources/$*
	touch $@

test: build
	go test ./...
	swift test

.PHONY: tools
tools:
	go generate -tags=tools .
