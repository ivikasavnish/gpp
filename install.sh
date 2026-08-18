#!/bin/bash
# Downloads and extracts the gpp compiler for Linux AMD64

set -e

RELEASE_URL="https://github.com/ivikasavnish/gpp/releases/latest/download/gpp-linux-amd64.tar.gz"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Downloading gpp compiler (Linux AMD64)..."
curl -L "$RELEASE_URL" -o /tmp/gpp-linux-amd64.tar.gz

echo "Extracting..."
tar -xzf /tmp/gpp-linux-amd64.tar.gz -C /tmp
mv /tmp/gpp-release "$SCRIPT_DIR/compiler"
rm /tmp/gpp-linux-amd64.tar.gz

chmod +x "$SCRIPT_DIR/gpp"
echo "Done. Run: ./gpp run test_decorator_syntax.go"
