package model

type ArgCreate struct {
	AppName string `json:"app_name"`
	TreeID  int    `json:"tree_id"`
	Env     string `json:"env" `
	Zone    string `json:"zone" `
	Version string `json:"version" `
	Comment string `json:"comment"`
	State   int    `json:"state"`
	From    int64  `json:"from,string" swaggertype:"string"`
}
