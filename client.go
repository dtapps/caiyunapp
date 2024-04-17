package caiyunapp

import (
	"go.dtapp.net/golog"
)

// Client 实例
type Client struct {
	token   string
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(token string) (*Client, error) {
	return &Client{token: token}, nil
}
