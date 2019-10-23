package api

import (
	"context"
	"valerian/library/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID discovery id.
const AppID = "admin.manager"

// NewClient .
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (PermitClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+AppID)
	if err != nil {
		return nil, err
	}
	return NewPermitClient(conn), nil
}