GOBIN=go
LDFLAGS="-extldflags '-static' -s -w"
GOBUILDFLAGS=-trimpath -race
OUT_BIN=mailgunostpiper

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -count 1 -race --coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Build for current arch / system
.PHONY: build
build:
	$(GOBIN) build -ldflags=$(LDFLAGS) $(GOBUILDFLAGS) -o build/$(OUT_BIN) main.go

# Clean builds
.PHONY: clean
clean:
	rm -Rf ./coverage.out
	rm -Rf ./build

# Build all (linux, darwin & windows)
.PHONY: build-all
build-all:
	make build-linux
	make build-darwin
	make build-windows

# Build for linux amd64
.PHONY: build-linux-fast
build-linux-fast:
	# Build for linux arch amd64
	GOOS=linux GOARCH=amd64 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-linux-amd64 main.go

# Build for linux archs 386, amd64, arm, arm64
.PHONY: build-linux
build-linux:
	# Build for linux arch 386
	GOOS=linux GOARCH=386 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-linux-386 main.go
	# Build for linux arch amd64
	GOOS=linux GOARCH=amd64 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-linux-amd64 main.go
	# Build for linux arch arm
	GOOS=linux GOARCH=arm $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-linux-arm main.go
	# Build for linux arch arm64
	GOOS=linux GOARCH=arm64 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-linux-arm64 main.go

# Build for darwin (osx) archs 386, amd64
.PHONY: build-darwin
build-darwin:
	# Build for darwin arch amd64
	GOOS=darwin GOARCH=amd64 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-darwin-amd64 main.go

# Build for windows archs 386, amd64
.PHONY: build-windows
build-windows:
	# Build for windows arch 386
	GOOS=windows GOARCH=386 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-win-386.exe main.go
	# Build for windows arch amd64
	GOOS=windows GOARCH=amd64 $(GOBIN) build $(BUILDARGS) -ldflags=$(LDFLAGS) -o build/$(OUT_BIN)-win-amd64.exe main.go

# Build all & zip
.PHONY: release
release:
	make clean
	make build-all
	cd build && find . -name "*" -exec zip {}.zip {} \;
	find ./build -type f ! -name '*.zip' -delete
