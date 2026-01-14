// Package main 是应用程序的主包
package main

import (
	"fmt"           // 导入格式化输出包
	"net/http"      // HTTP包
	"runtime/debug" // 运行时调试包，用于获取堆栈跟踪
)

// serverError 方法用于处理服务器内部错误（500）
func (app *application) serverError(w http.ResponseWriter, err error) {
	// 生成错误信息和堆栈跟踪
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// 将错误信息写入错误日志
	app.errorLog.Output(2, trace)

	// 向客户端返回500错误
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError 方法用于处理客户端错误
func (app *application) clientError(w http.ResponseWriter, status int) {
	// 根据指定的状态码向客户端返回错误响应
	http.Error(w, http.StatusText(status), status)
}

// notFound 方法用于处理资源未找到错误（404）
func (app *application) notFound(w http.ResponseWriter) {
	// 调用clientError方法返回404错误
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// 从模板缓存中获取指定名称的模板
	ts, ok := app.templateCache[page]
	if !ok {
		// 如果模板不存在，返回服务器错误
		app.serverError(w, fmt.Errorf("The template %s does not exist", page))
		return
	}

	// 执行模板渲染
	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		// 如果渲染过程中发生错误，返回服务器错误
		app.serverError(w, err)
	}
}
