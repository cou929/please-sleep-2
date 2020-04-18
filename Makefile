SRCS := $(shell find . -type f -name '*.go')
POSTS := $(shell find ./post -type f)
VIEWS := $(shell find ./view -type f)

all: test dist

dist: $(POSTS) $(SRCS) $(VIEWS)
	make clean
	make run
	make asset

.PHONY: clean
clean:
	rm -rf dist

.PHONY: run
run:
	mkdir -p dist
	go run .

.PHONY: asset
asset:
	mkdir -p dist
	cp -r view/static/* dist/
	cp -r post/images dist/

.PHONY: test
test:
	go test -v ./...

.PHONY: watch
watch:
	while true; do \
		make dist --silent; \
		sleep 1; \
	done
