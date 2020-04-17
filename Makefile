BINNAME=ps2
DISTDIR=dist
SRCS := $(shell find . -type f -name '*.go')
POSTS := $(shell find ./post -type f)

all: test clean dist

$(BINNAME): $(SRCS)
	go build -o $(BINNAME)

dist: $(BINNAME) $(POSTS)
	make clean
	./$(BINNAME)
	make asset

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf $(DISTDIR)

.PHONY: distinit
distinit: clean
	mkdir -p dist

.PHONY: asset
asset: distinit
	cp -r view/static/* dist/
	cp -r post/images dist/

.PHONY: watch
watch:
	while true; do \
		make dist --silent; \
		sleep 3; \
	done
