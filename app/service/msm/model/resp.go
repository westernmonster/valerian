package model

import "container/list"

// Code ver and message.
type ErrCode struct {
	Ver  int64
	Code int
	Msg  string
}

// Codes all codes local map cache.
type ErrCodes struct {
	Ver  int64
	MD5  string
	Code map[int]string
}

type ErrCodesLangs struct {
	Ver  int64
	MD5  string
	Code map[int]map[string]string
}

type ErrCodeLangs struct {
	Ver  int64
	Code int
	Msg  map[string]string
}

// Version list and map.
type Version struct {
	List *list.List
	Map  map[int64]*list.Element
}
