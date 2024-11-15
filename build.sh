#!/bin/bash

targets=(
  "linux/amd64"
  "windows/amd64"
  "darwin/amd64"
  "darwin/arm64"
)

# Build the binaries
for target in "${targets[@]}"; do
  os=$(echo "$target" | cut -d'/' -f1)
  arch=$(echo "$target" | cut -d'/' -f2)
  output="dist/network_test-${os}-${arch}"

  if [ "$os" = "windows" ]; then
    output="${output}.exe"
  fi

  echo "Building for $os/$arch..."
  env GOOS="$os" GOARCH="$arch" go build -o "$output" ./cmd/network_test
done

echo "Build complete!"

