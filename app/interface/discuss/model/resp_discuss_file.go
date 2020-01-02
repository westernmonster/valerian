package model

type DiscussionFileResp struct {
	ID        int64  `json:"id,string" swaggertype:"string"`
	FileName  string `json:"file_name"` // FileName 文件名
	FileURL   string `json:"file_url"`  // FileURL 文件地址
	PdfURL    string `json:"pdf_url"`   // PDFURL PDF文件地址
	FileType  string `json:"file_type"` // FileType 文件类型
	Seq       int32  `json:"seq"`       // Seq 文件顺序
	CreatedAt int64  `json:"created_at"`
}
