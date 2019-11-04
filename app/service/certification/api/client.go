package api

import (
	"context"
	"time"

	"valerian/library/net/rpc/warden"
	xtime "valerian/library/time"

	"google.golang.org/grpc"
)

// AppID unique app id for service discovery
const AppID = "service.certification"

// NewClient new member grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (CertificationClient, error) {
	cfg.Timeout = xtime.Duration(time.Second * 3)
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+AppID)
	if err != nil {
		return nil, err
	}
	return NewCertificationClient(conn), nil
}
