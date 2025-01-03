package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// 定义一个结构体，用于描述 Swagger 文档中的标签列表和分页信息
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单个文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取多个文章
// @Produce  json
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary 新增文章
// @Produce  json
// @Param tag_id body int false "标签ID"
// @Param title body string true "标题" minlength(3) maxlength(100)
// @Param desc body string false "简述" maxlength(255)
// @Param content body string false "内容" maxlength(4294967295)
// @Param cover_image_url body string false "封面图" maxlength(255)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章ID"
// @Param tag_id body int false "标签ID"
// @Param title body string false "标题" minlength(3) maxlength(100)
// @Param desc body string false "简述" maxlength(255)
// @Param content body string false "内容" maxlength(4294967295)
// @Param cover_image_url body string false "封面图" maxlength(255)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
