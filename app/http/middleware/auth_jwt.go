package middleware

import (
	"github.com/gin-gonic/gin"
	"go-hub/pkg/jwt"
	"go-hub/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(ctx)
		if err != nil {
			response.ErrAuth(ctx, err.Error())
			return
		}

		if claims.UserID == 0 {
			response.BadRequest(ctx, "用户不存在")
			return
		}

		ctx.Set("current_user_id", claims.UserID)
		ctx.Next()
	}
}
