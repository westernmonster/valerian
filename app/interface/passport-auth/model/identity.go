package model

type IdentityInfo struct {
	Aid     int64  `json:"aid,string" swaggertype:"string"`
	Csrf    string `json:"csrf"`
	Expires int    `json:"expires"`
}
