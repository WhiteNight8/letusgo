// Package models 包含应用程序的数据模型和数据库操作
package models

import (
	"errors" // 导入标准错误包
)

// ErrNoRecord 是一个自定义错误，用于表示未找到匹配记录的情况
var ErrNoRecord = errors.New("models: no matching record found")
