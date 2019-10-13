package model

type ArgProcessInvite struct {
	// 请求的ID
	ID int64 `json:"id,string" swaggertype:"string"`

	Result bool `json:"result"`
}

func (p *ArgProcessInvite) Validate() (err error) {
	return
}
