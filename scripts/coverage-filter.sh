#!/bin/bash

# Script to filter coverage based on .covignore file
# Usage: ./scripts/coverage-filter.sh

set -e

COVERAGE_FILE="coverage.out"
COVIGNORE_FILE=".covignore"
FILTERED_COVERAGE_FILE="coverage-filtered.out"

if [[ ! -f "$COVERAGE_FILE" ]]; then
    echo "Coverage file $COVERAGE_FILE not found"
    exit 1
fi

# Start with the original coverage file
cp "$COVERAGE_FILE" "$FILTERED_COVERAGE_FILE"

# If .covignore exists, filter out patterns
if [[ -f "$COVIGNORE_FILE" ]]; then
    echo "Filtering coverage using $COVIGNORE_FILE"
    
    # Read each pattern from .covignore and remove matching lines from coverage
    while IFS= read -r pattern; do
        # Skip empty lines and comments
        if [[ -z "$pattern" ]] || [[ "$pattern" == \#* ]]; then
            continue
        fi
        
        # Remove leading slash if present for pattern matching
        pattern="${pattern#/}"
        
        echo "Excluding pattern: $pattern"
        
        # Use grep to exclude lines matching the pattern
        # The pattern should match the file path in coverage.out
        grep -v "$pattern" "$FILTERED_COVERAGE_FILE" > temp_coverage.out || true
        mv temp_coverage.out "$FILTERED_COVERAGE_FILE"
        
    done < "$COVIGNORE_FILE"
    
    echo "Coverage filtering complete"
    echo "Original coverage entries: $(wc -l < "$COVERAGE_FILE")"
    echo "Filtered coverage entries: $(wc -l < "$FILTERED_COVERAGE_FILE")"
else
    echo "No $COVIGNORE_FILE found, using original coverage"
fi

# Calculate and display coverage percentage
if [[ -s "$FILTERED_COVERAGE_FILE" ]]; then
    echo "Calculating coverage from filtered results..."
    COVERAGE=$(go tool cover -func="$FILTERED_COVERAGE_FILE" | grep '^total:' | awk '{print $3}' | sed 's/%//')
    echo "Filtered coverage: ${COVERAGE}%"
else
    echo "Warning: Filtered coverage file is empty"
    COVERAGE="0.0"
fi

# Export coverage for use in CI
echo "FILTERED_COVERAGE=${COVERAGE}" >> "$GITHUB_ENV" || true
echo "Filtered coverage is ${COVERAGE}%"