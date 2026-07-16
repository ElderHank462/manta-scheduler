VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || git rev-parse --short HEAD)

.PHONY: build
build:
	go build -ldflags="-X manta/cmd.version=$(VERSION)" .

.PHONY: clean
clean:
	rm -f manta

.PHONY: run
run: build
	./manta
