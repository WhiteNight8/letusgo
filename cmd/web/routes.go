// Package main 是应用程序的主包
package main

import (
	"net/http" // 导入HTTP包
)

// routes 方法用于配置应用程序的路由
func (app *application) routes() *http.ServeMux {
	// 创建一个新的ServeMux实例
	mux := http.NewServeMux()

	// 配置静态文件服务
	// http.Dir指定静态文件目录
	// http.StripPrefix用于去除URL路径中的/static前缀
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// 配置主页路由
	mux.HandleFunc("/", app.home)
	// 配置查看代码片段路由
	mux.HandleFunc("/letusgo/view", app.letusgoView)
	// 配置创建代码片段路由
	mux.HandleFunc("/letusgo/create", app.letusgoCreate)

	// 返回配置好的ServeMux
	return mux
}
