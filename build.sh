#!/bin/sh

echo "Fetching tag and tidying..."

TAG=$(git describe --tags $(git rev-list --tags --max-count=1))

go mod tidy

for OS in darwin linux; do
    echo "Building $OS executable."
    rm -rf ./demon-$OS-*
    GOOS=$OS GOARCH=amd64 go build -o demon-$OS-$TAG .
done

echo "Done!"
