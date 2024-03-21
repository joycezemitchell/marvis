.PHONY: build build-windows build-debian arm

GOOS ?= linux
GOARCH ?= arm
GOARM ?= 6
OUTFILE ?= marvis

build:
	@echo "Building Marvis..."
	cd cmd && GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -o $(OUTFILE) main.go && cp $(OUTFILE) ../build
	cp config/config.yaml build/config
	@echo "Build complete"

build-windows:
	@echo "Building Marvis for Windows..."
	$(MAKE) build GOOS=windows GOARCH=amd64 OUTFILE=marvis.exe
	@echo "Build complete for Windows"

build-debian:
	@echo "Building Marvis for Debian..."
	$(MAKE) build GOOS=linux GOARCH=amd64
	@echo "Build complete for Debian"

build-arm:
	@echo "Building Marvis for Raspberry Pi..."
	$(MAKE) build GOOS=linux GOARCH=arm GOARM=6
	@echo "Build complete for Raspberry Pi"		

