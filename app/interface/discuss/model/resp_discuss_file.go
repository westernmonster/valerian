package model

type DiscussionFileResp struct {
	ID        int64  `json:"id,string" swaggertype:"string"`
	FileName  string `json:"file_name"` // FileName 文件名
	FileURL   string `json:"file_url"`  // FileURL 文件地址
	Seq       int    `json:"seq"`       // Seq 文件顺序
	CreatedAt int64  `json:"created_at"`
}
