#!/bin/bash

set -e

echo "=========================================="
echo "DocsGo Cross-Platform Build Script"
echo "=========================================="
echo ""

# Create bin directory
mkdir -p bin

# Clean old builds
echo "Cleaning old builds..."
rm -f bin/docs-go-*
echo ""

# Get version from git tag or use default
VERSION=$(git describe --tags --always 2>/dev/null || echo "v0.1")
echo "Building version: $VERSION"
echo ""

# Build flags
LDFLAGS="-s -w"

# Function to build for a specific platform
build_platform() {
    local os=$1
    local arch=$2
    local output=$3
    
    echo "=========================================="
    echo "Building for $os ($arch)..."
    echo "=========================================="
    
    GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -o "$output"
    
    if [ $? -eq 0 ]; then
        echo "✓ $output"
    else
        echo "✗ Failed to build for $os $arch"
        return 1
    fi
    echo ""
}

# Build for all platforms
build_platform "windows" "amd64" "bin/docs-go-windows-amd64.exe"
build_platform "windows" "arm64" "bin/docs-go-windows-arm64.exe"
build_platform "linux" "amd64" "bin/docs-go-linux-amd64"
build_platform "linux" "arm64" "bin/docs-go-linux-arm64"
build_platform "darwin" "amd64" "bin/docs-go-darwin-amd64"
build_platform "darwin" "arm64" "bin/docs-go-darwin-arm64"

echo "=========================================="
echo "Build Summary"
echo "=========================================="
ls -lh bin/docs-go-* 2>/dev/null || echo "No builds found"
echo ""
echo "Total builds completed successfully!"
