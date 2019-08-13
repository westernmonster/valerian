package model

type ArgTopicFollow struct {
	TopicID       int64  `json:"topic_id,string" swaggertype:"string"`
	Reason        string `json:"reason"`
	AllowViewCert bool   `json:"allow_view_cert"`
}

func (p *ArgTopicFollow) Validate() (err error) {
	return
}
