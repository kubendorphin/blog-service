package model

import (
	"blog-service/pkg/app"
	"gorm.io/gorm"
)

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

func (t *Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	db = db.Model(&Tag{})
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Where("id =? and is_del =?", t.ID, 0).Updates(&t).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.ID, 0).Delete(&t).Error
}

// 补充定义标签的方法
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name =?", t.Name)
	}
	return tags, err
}
