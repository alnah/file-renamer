# File Renamer CLI
BINARY_NAME=file-renamer
GO_FILES=*.go

# Build for current platform
build:
	go build -o $(BINARY_NAME) $(GO_FILES)

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 $(GO_FILES)
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe $(GO_FILES)
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64 $(GO_FILES)

# Install to $GOPATH/bin
install:
	go install

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)*

# Format code
fmt:
	go fmt ./...

# Run with example
run:
	go run $(GO_FILES) -help

# Quick test with dry-run
test:
	@echo "Creating test files..."
	@touch this_pattern_1.txt this_pattern_2.txt this_pattern_3.doc
	@echo "Running dry-run test..."
	go run $(GO_FILES) -from 'this_pattern_*.txt' -to 'new_pattern_*.txt' -dry-run
	@echo "Cleaning up test files..."
	@rm -f this_pattern_*.txt this_pattern_*.doc

.PHONY: build build-all install clean fmt run test
