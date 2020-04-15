BINNAME=ps2
DISTDIR=dist
SRCS := $(shell find . -type f -name '*.go')

all: test clean run asset

$(BINNAME): $(SRCS)
	go build -o $(BINNAME)

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf $(BINNAME)
	rm -rf $(DISTDIR)

.PHONY: distinit
distinit: clean
	mkdir -p dist

.PHONY: asset
asset: distinit
	cp -r view/static/* dist/
	cp -r post/images dist/
