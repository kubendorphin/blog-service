package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// TableName 方法应该属于 ArticleTag 类型，并且接收者应该是 ArticleTag 类型的值
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
