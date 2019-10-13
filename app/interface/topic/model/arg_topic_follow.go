package model

type ArgTopicFollow struct {
	TopicID       int64  `json:"topic_id,string" swaggertype:"string"`
	Reason        string `json:"reason"`
	AllowViewCert bool   `json:"allow_view_cert"`
}

func (p *ArgTopicFollow) Validate() (err error) {
	return
}

type ArgAuditFollow struct {
	// 请求的ID
	ID int64 `json:"id,string" swaggertype:"string"`

	Approve bool `json:"approve"`
}

func (p *ArgAuditFollow) Validate() (err error) {
	return
}
