# Simple Makefile to generate protobufs using buf
# Usage:
#   make proto        Generate protobuf code via buf
#   make tools        Check that required tools are installed
#   make help         Show this help

.PHONY: proto tools help

BUF := $(shell command -v buf 2>/dev/null)

help:
	@echo "Targets:"
	@echo "  proto  - Run 'buf generate' to (re)generate protobuf code"
	@echo "  tools  - Check required tools are installed"

# Ensure buf is available and run generation.
proto: tools
	@echo "Running buf generate..."
	buf generate

# Check for required tools. Fails with instructions if missing.
tools:
ifndef BUF
	@echo "Error: 'buf' not found in PATH."
	@echo "Install buf: https://buf.build/docs/installation"
	@exit 1
else
	@echo "buf found: $(BUF)"
endif
