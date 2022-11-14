package jwt

import "github.com/gin-gonic/gin"

func CurrentUID(c *gin.Context) int64 {
	return c.GetInt64("current_user_id")
}
