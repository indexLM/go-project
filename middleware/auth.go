package middleware

import (
	"github.com/gin-gonic/gin"
	response "go-project/model/resopnse"
	"go-project/utils"
)

func Auth(context *gin.Context) {
	header := context.GetHeader("Token")
	if header == "" {
		//封装通用json返回
		response.FailWithDetailed(response.AUTH, nil, "无效的请求头", context)
		//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
		context.Abort()
		return
	}
	err := utils.JwtVerification(header, context)
	if err != nil {
		//封装通用json返回
		response.FailWithDetailed(response.AUTH, nil, err.Error(), context)
		//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
		context.Abort()
		return
	}
	context.Next()
}
