package app

import (
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 定义一个响应数据结构体
type Response struct {
	Ctx *gin.Context
}

// 定义一个分页器结构体
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

// 定义了一个名为 NewResponse 的函数。
// 该函数接受一个 *gin.Context 类型的参数，通过&Response{Ctx: ctx}创建并返回一个指向 Response 结构体的指针。
// 这个函数的作用是创建一个新的 Response 实例，并将传入的 *gin.Context 赋值给 Response 结构体的 Ctx 字段。
// 这样，通过调用 NewResponse 函数，我们可以方便地创建一个 Response 实例，并在后续的处理中使用它。
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// 定义一个函数，用于获取分页参数
// 接受一个 *gin.Context 类型的参数，检查传入的 data 参数是否为 nil 。如果是 nil ，则将其设置为一个空的 gin.H。
// 然后，通过 r.Ctx.JSON(200, data) ，使用与 r 关联的 gin.Context （即 r.Ctx ）以 200 状态码将 data 以 JSON 格式返回给客户端。
// 这个方法用于将给定的数据以成功状态（200）的 JSON 响应发送给请求方，如果数据为 nil 则使用一个默认的空数据。
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(200, data)
}

// 这个方法的作用是构建一个包含列表数据、分页信息并以 JSON 格式进行响应的输出。
// 首先，通过 r.Ctx.JSON(200,...) 准备以 200 状态码发送一个 JSON 响应
// 在构建的 JSON 数据中：
// "list" 字段的值为传入的 list 接口参数，表示要返回的列表数据；
// "pager" 字段的值为一个 Pager 结构体，其中包含了分页信息，包括当前页码、每页大小和总记录数。
// 这个方法用于构建一个包含列表数据和分页信息的响应，以便在客户端进行展示。
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(200, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

// 这个方法的作用是构建一个包含错误信息的 JSON 响应，并将其发送给客户端。
// func (r *Response) ToErrorResponse(err *errcode.Error)：这是一个为 *Response 类型定义的方法，
// 接收一个指向 errcode.Error 类型的指针 err 。
// response := gin.H{"code": err.Code(), "msg": err.Msg()}：创建一个 gin.H 类型的映射（类似于字典）response ，
// 其中包含两个键值对："code" 对应着错误的代码（通过 err.Code() 获取），"msg" 对应着错误的消息（通过 err.Msg() 获取）。
// details := err.Details()：获取错误的详细信息，通过 err.Details() 获取。
// if len(details) > 0：如果错误的详细信息长度大于 0，说明有额外的错误细节。
// response["details"] = details：将错误的详细信息添加到 response 映射中，键为 "details" 。
// r.Ctx.JSON(err.StatusCode(), response)：使用 r.Ctx.JSON() 方法将 response 映射以 JSON 格式发送给客户端，
// 状态码为 err.StatusCode() 获取的状态码。
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
