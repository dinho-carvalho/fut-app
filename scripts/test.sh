#!/usr/bin/env bash

set -e

cleanup() {
  rm -f coverage.txt coverage.tmp profile.out gocov.json
}

trap cleanup EXIT

cleanup
echo "mode: set" > coverage.tmp

for d in $(go list ./... | grep -v vendor); do
  go test -coverprofile=profile.out -covermode=set $d

  if [ -f profile.out ]; then
    tail -n +2 profile.out >> coverage.tmp
    rm profile.out
  fi
done

grep -v -E -f .covignore coverage.tmp > coverage.filtered.out
mv coverage.filtered.out coverage.tmp

go run github.com/axw/gocov/gocov@latest convert coverage.tmp > gocov.json
go run github.com/matm/gocov-html/cmd/gocov-html@latest -t kit gocov.json > gocov.html

open gocov.html
