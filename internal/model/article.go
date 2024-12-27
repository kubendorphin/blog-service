package model

import "blog-service/pkg/app"

type Article struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// 定义一个结构体，用于描述 Swagger 文档中的标签列表和分页信息
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// TableName 方法应该属于 ArticleTag 类型，并且接收者应该是 ArticleTag 类型的值
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
