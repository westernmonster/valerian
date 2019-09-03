package api

import (
	"context"

	"google.golang.org/grpc"

	"valerian/library/net/rpc/warden"
)

// AppID unique app id for service discovery
const AppID = "service.identify"

// NewClient new identify grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (IdentifyClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+AppID)
	if err != nil {
		return nil, err
	}
	return NewIdentifyClient(conn), nil
}
