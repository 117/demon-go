version=0.0.1

GOOS=linux GOARCH=amd64 go build -o ../builds/demon-linux-$version ../main.go
GOOS=darwin GOARCH=amd64 go build -o ../builds/demon-darwin-$version ../main.go