package model

type AppListResp struct {
	Total    int32      `json:"total"`
	PageSize int32      `json:"page_size"`
	Page     int32      `json:"page"`
	Items    []*AppItem `json:"items"`
}

type AppItem struct {
	ID        int64  `json:"id,string" swaggertype:"string"`
	Name      string `json:"name"`
	Token     string `json:"token"`
	Env       string `json:"env"`
	Zone      string `json:"zone"`
	Platform  string `json:"platform"`
	TreeID    int32  `json:"tree_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
