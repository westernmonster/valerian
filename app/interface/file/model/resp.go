package model

type ConfigStruct struct {
	Expiration string          `json:"expiration"`
	Conditions [][]interface{} `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire,string" swaggertype:"string"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
	Key         string `json:"key"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callback_url"`
	CallbackBody     string `json:"callback_body"`
	CallbackBodyType string `json:"callback_body_type"`
}

type STSResp struct {
	AccessKeySecret string `json:"access_key_secret"`
	Expiration      string `json:"expiration"`
	AccessKeyId     string `json:"access_key_id"`
	SecurityToken   string `json:"security_token"`
}
