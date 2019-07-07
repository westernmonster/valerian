#! /bin/sh
# proto.sh
gopath=$GOPATH/src
gogopath=$GOPATH/src/valerian/vendor/github.com/gogo/protobuf
protoc --gofast_out=. --proto_path=$gopath:$gogopath:. *.proto
protoc-go-inject-tag -input=./auth.pb.go



