package model

type AuthReply struct {
	Login bool  `json:"login"`
	Aid   int64 `json:"Aid,string" swaggertype:"string"`
}
