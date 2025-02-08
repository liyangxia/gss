# Copyright 2025 liyangxia.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

# Project name
PROJECT_NAME := gss

# Output directory
OUTPUT_DIR := dist

# Supported operating systems and architectures
OS := linux darwin
ARCH := amd64 arm64

# Default target
.PHONY: default
default: build

# Clean the output directory
.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

# Build binary for the current system environment
.PHONY: build
build:
	@echo "Building for current system environment..."; \
	go build -o $(OUTPUT_DIR)/$(PROJECT_NAME) ./src/

# Build binaries for all supported platforms and architectures
.PHONY: build-all
build-all:
	@for os in $(OS); do \
		for arch in $(ARCH); do \
			echo "Building $$os/$$arch..."; \
			GOOS=$$os GOARCH=$$arch go build -o $(OUTPUT_DIR)/$(PROJECT_NAME)-$$os-$$arch ./src/; \
		done; \
	done
