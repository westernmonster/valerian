package models

type Actor struct {
	Avatar       string `json:"avatar"`
	Gender       int    `json:"gender"`
	Introduction string `json:"introduction"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	IsFollowed   bool   `json:"is_followed"`
	IsFollowing  bool   `json:"is_following"`
	UserType     string `json:"user_type"`
	ID           int64  `json:"id,string" swaggertype:"string"`
}

type Activity struct {
	ID         int64       `json:"id,string" swaggertype:"string"`
	ActionText string      `json:"action_text"`
	CreatedAt  int64       `json:"created_at"`
	Verb       string      `json:"verb"`
	Type       string      `json:"type"`
	Target     interface{} `json:"target"`
	Actor      Actor       `json:"actor"`
}
