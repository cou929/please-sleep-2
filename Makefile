GEN_SRCS := $(shell find cmd/gen/ -type f -name '*.go')
POSTS := $(shell find ./post -type f)
VIEWS := $(shell find ./view -type f)

all: test dist

dist: $(POSTS) $(GEN_SRCS) $(VIEWS)
	make clean
	make run
	make asset

gen: $(GEN_SRCS)
	go build -o gen ./cmd/gen

.PHONY: clean
clean:
	rm -rf dist
	rm -f gen

.PHONY: run
run: gen
	mkdir -p dist
	./gen

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

.PHONY: local-server
local-server:
	go run ./tools/localsvr/ ./dist/

.PHONY: dev
dev:
	make watch &
	make local-server
