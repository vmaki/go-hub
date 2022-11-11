package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-hub/common/helpers"
	"go-hub/pkg/logger"
	"go.uber.org/zap"
	"io"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)                 // c.Request.Body 是一个 buffer 对象，只能读取一次
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody)) // 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
		}

		// 设置开始时间
		start := time.Now()

		c.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start)
		status := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", status),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", helpers.ClientIP(c.Request)),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			logFields = append(logFields, zap.String("Request Body", string(requestBody))) // 请求的内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))    // 响应的内容
		}

		if status > 400 && status <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(status), logFields...)
		} else if status >= 500 && status <= 599 {
			logger.Error("HTTP Error "+cast.ToString(status), logFields...)
		}
	}
}
