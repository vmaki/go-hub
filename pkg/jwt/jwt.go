// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"go-hub/common/helper"
	"go-hub/config"
	"go-hub/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	SignKey    []byte
	MaxRefresh time.Duration
}

// CustomClaims 自定义载荷
type CustomClaims struct {
	UserID       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口并且继承了 Valid() 方法
	jwtPkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.Cfg.JWT.SignKey),
		MaxRefresh: time.Duration(config.Cfg.JWT.MaxRefresh) * time.Minute,
	}
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	expire := time.Duration(config.Cfg.JWT.ExpireTime) * time.Minute
	return helper.NowTime().Add(expire).Unix()
}

// CreateToken 创建 Token
func (jwt *JWT) CreateToken(userID int64, userName string) string {
	expireAtTime := jwt.expireAtTime()

	claims := CustomClaims{
		userID,
		userName,
		expireAtTime,
		jwtPkg.StandardClaims{
			NotBefore: helper.NowTime().Unix(),
			IssuedAt:  helper.NowTime().Unix(),
			ExpiresAt: expireAtTime,
			Issuer:    config.Cfg.Application.Name,
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.new(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

func (jwt *JWT) new(claims CustomClaims) (string, error) {
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// getTokenFromHeader 从 header 获取 token
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}

// parseTokenString 解析 token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// ParserToken 解析 token
func (jwt *JWT) ParserToken(c *gin.Context) (*CustomClaims, error) {
	// 1. 从 header 获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 2. 解析Token
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtPkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtPkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 刷新 token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 1. 从 header 获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 解析Token
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 3. 解析 CustomClaims 的数据
	claims := token.Claims.(*CustomClaims)

	// 4. 检查是否过了『最大允许刷新的时间』
	last := helper.NowTime().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > last {
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.new(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}
