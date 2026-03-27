#!/bin/bash

# GKA Cross-Platform Build Script
# This script builds the GKA binary for macOS, Linux, and Windows.

APP_NAME="gka"
BUILD_DIR="build"

echo "Creating build directory..."
mkdir -p $BUILD_DIR

# 1. macOS (Intel)
echo "Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_darwin_amd64 .

# 2. macOS (Apple Silicon)
echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/${APP_NAME}_darwin_arm64 .

# 3. Linux (amd64)
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_linux_amd64 .

# 4. Windows (amd64)
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_windows_amd64.exe .

echo "Done! Binaries are available in the /$BUILD_DIR folder."
ls -l $BUILD_DIR
