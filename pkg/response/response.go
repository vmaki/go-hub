package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
func defaultMessage(defaultMsg string, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}

	return defaultMsg
}

func Success(ctx *gin.Context, msg ...string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": defaultMessage("success", msg...),
	})
}

func JSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func Abort404(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"code":    404,
		"message": defaultMessage("该接口不存在，请确定请求正确", msg...),
	})
}

func Error(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": defaultMessage("服务器内部错误，请稍后再试", msg...),
	})
}

func BadRequest(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code":    4001,
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确", msg...),
	})
}

func ErrAuth(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code":    4101,
		"message": defaultMessage("请重新授权", msg...),
	})
}
