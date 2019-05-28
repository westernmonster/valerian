package model

type AuthReply struct {
	Login     bool   `json:"login"`
	Aid       int64  `json:"aid,string" swaggertype:"string"`
	Csrf      string `json:"csrf"`
	ExpiresAt int64  `json:"expires_at"`
}
