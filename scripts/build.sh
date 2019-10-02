repo=$(dirname "$0")

if [ -z $1 ]
then
    echo "no version tag provided"
else
    rm -rf $repo/../builds
    mkdir $repo/../builds
    GOOS=linux GOARCH=amd64 go build -o $repo/../builds/demon-linux-$1 ../main.go
    GOOS=darwin GOARCH=amd64 go build -o $repo/../builds/demon-darwin-$1 ../main.go
fi



