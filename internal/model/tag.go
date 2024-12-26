package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// 确保 TableName 方法的接收者名称和结构体名称一致，这里是 Tag
func (t *Tag) TableName() string {
	return "blog_tag"
}
