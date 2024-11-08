# Variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BINARY_NAME ?= melodica-$(GOOS)-$(GOARCH)
DIST_DIR ?= dist
CMD_DIR ?= cmd/melodica/main.go
PLAYLIST ?= playlist.txt

# Default Go version (can be overridden)
GO_VERSION ?= 1.23.2

# Set Go build parameters
BUILD_CMD = GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(DIST_DIR)/$(BINARY_NAME) $(CMD_DIR)

.PHONY: all build test clean release run deps

all: build

# Install dependencies
deps:
	@echo "Installing dependencies..."
ifeq ($(GOOS),linux)
	sudo apt update
	sudo apt install -y libasound2-dev
endif

# Build the application
build: deps
	@echo "Building application for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(DIST_DIR)
	$(BUILD_CMD)

# Run the application with a specified playlist
run: build
	@echo "Running application with playlist: $(PLAYLIST)..."
	./$(DIST_DIR)/$(BINARY_NAME) $(PLAYLIST)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Package the binary and playlist into a zip file for release
release: build
	@echo "Creating zip archive for release..."
	cp $(PLAYLIST) $(DIST_DIR)/
	cd $(DIST_DIR) && zip "$(BINARY_NAME).zip" "$(BINARY_NAME)" $(PLAYLIST)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(DIST_DIR)
