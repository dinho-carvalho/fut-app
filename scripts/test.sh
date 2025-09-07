#!/usr/bin/env bash

set -e

echo "🧪 Running tests and generating coverage dashboard..."

cleanup() {
  rm -f coverage.txt coverage.tmp profile.out gocov.json
}

trap cleanup EXIT

cleanup
echo "mode: set" > coverage.tmp

# Test the packages we care about (same as pipeline)
PACKAGES="./internal/errors/... ./pkg/logger/... ./internal/handlers/... ./internal/domain/... ./internal/database/repositories/... ./internal/usecase/..."

echo "📦 Testing packages: $PACKAGES"

for pkg in $PACKAGES; do
  echo "Testing $pkg..."
  go test -coverprofile=profile.out -covermode=set $pkg

  if [ -f profile.out ]; then
    tail -n +2 profile.out >> coverage.tmp
    rm profile.out
  fi
done

echo "🔍 Applying coverage filters..."

# Apply filtering using our patterns file if it exists
if [ -f .covignore-patterns ]; then
  echo "Using .covignore-patterns for filtering"
  grep -v -f .covignore-patterns coverage.tmp > coverage.filtered.out || cp coverage.tmp coverage.filtered.out
  mv coverage.filtered.out coverage.tmp
else
  echo "No .covignore-patterns found, using all coverage"
fi

echo "📊 Generating HTML dashboard..."

# Generate coverage dashboard
go run github.com/axw/gocov/gocov@latest convert coverage.tmp > gocov.json
go run github.com/matm/gocov-html/cmd/gocov-html@latest -t kit gocov.json > gocov.html

# Calculate and display coverage percentage
COVERAGE=$(go tool cover -func=coverage.tmp | grep '^total:' | awk '{print $3}' | sed 's/%//')
echo "✅ Total coverage: ${COVERAGE}%"

echo "🌐 Opening coverage dashboard in browser..."
open gocov.html

echo "📁 Coverage files generated:"
echo "  - gocov.html (interactive dashboard)"
echo "  - coverage.tmp (coverage profile)"
echo "  - gocov.json (JSON coverage data)"
