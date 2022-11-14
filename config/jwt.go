package config

type JWTConfig struct {
	SignKey    string
	ExpireTime int64
	MaxRefresh int64
}
