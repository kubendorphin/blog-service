package model

import "blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// 定义一个结构体，用于描述 Swagger 文档中的标签列表和分页信息
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// 确保 TableName 方法的接收者名称和结构体名称一致，这里是 Tag
func (t *Tag) TableName() string {
	return "blog_tag"
}
