#! /bin/sh
# proto.sh
gopath=$GOPATH/src
gogopath=$GOPATH/src/valerian/vendor/github.com/gogo/protobuf
protoc --gogofast_out=plugins=grpc:. --proto_path=$gopath:$gogopath:. *.proto





