package mq

import (
	xtime "valerian/library/time"
)

type Config struct {
	Nodes          []string
	ClusterID      string
	ClientID       string
	ConnectTimeout xtime.Duration
	AckTimeout     xtime.Duration
	PingInterval   int
	PingMaxOut     int
}
