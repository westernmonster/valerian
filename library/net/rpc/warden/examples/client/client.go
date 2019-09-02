package main

import (
	"context"
	"flag"
	"fmt"

	"valerian/library/ecode"
	"valerian/library/net/metadata"
	"valerian/library/net/rpc/warden"
	pb "valerian/library/net/rpc/warden/proto/testproto"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
)

// usage: ./client -grpc.target=test.service=127.0.0.1:8080
func main() {
	flag.Parse()
	conn, err := warden.NewConn("discovery://d/test.service?t=t&cluster=asdasd")
	if err != nil {
		panic(err)
	}
	cli := pb.NewGreeterClient(conn)
	normalCall(cli)
	errDetailCall(cli)
}

func normalCall(cli pb.GreeterClient) {
	ctx := metadata.NewContext(context.Background(), metadata.MD{
		metadata.Color: "red",
	})
	reply, err := cli.SayHello(ctx, &pb.HelloRequest{Name: "tom", Age: 23})
	if err != nil {
		panic(err)
	}
	fmt.Println("get reply:", *reply)
}

func errDetailCall(cli pb.GreeterClient) {
	_, err := cli.SayHello(context.Background(), &pb.HelloRequest{Name: "err_detail_test", Age: 12})
	if err != nil {
		any := errors.Cause(err).(ecode.Codes).Details()[0].(*any.Any)
		var re pb.HelloReply
		err := ptypes.UnmarshalAny(any, &re)
		if err == nil {
			fmt.Printf("cli.SayHello get error detail!detail:=%v", re)
		}
		return
	}
}
