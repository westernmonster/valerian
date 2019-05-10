package dao

import "valerian/library/cache/memcache"

type Dao struct {
	// c      *conf.Config
	authMC *memcache.Pool
}
