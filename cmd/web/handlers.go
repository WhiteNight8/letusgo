// Package main 是应用程序的主包
package main

import (
	"errors" // 导入错误处理包
	"fmt"    // 导入格式化输出包
	// HTML模板包（当前已注释）
	"net/http" // HTTP包
	"strconv"  // 字符串转换包

	"letgo.snippetbox/internal/models" // 内部数据模型包
)

// home 处理函数用于处理主页请求
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 检查请求路径是否为根路径
	if r.URL.Path != "/" {
		app.notFound(w) // 路径错误时返回404
		return
	}

	// 获取最近创建的10个代码片段
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err) // 数据库查询失败时返回500
		return
	}

	// 渲染主页模板
	app.render(w, http.StatusOK, "home.html", &templateData{
		Snippets: snippets,
	})
}

// letusgoView 处理函数用于查看单个代码片段
func (app *application) letusgoView(w http.ResponseWriter, r *http.Request) {
	// 从URL查询参数中获取id值
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // id无效时返回404
		return
	}

	// 根据id获取对应的代码片段
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w) // 未找到记录时返回404
		} else {
			app.serverError(w, err) // 其他错误返回500
		}
		return
	}

	// 渲染查看代码片段模板
	app.render(w, http.StatusOK, "view.html", &templateData{
		Snippet: snippet,
	})
}

// letusgoCreate 处理函数用于创建新的代码片段
func (app *application) letusgoCreate(w http.ResponseWriter, r *http.Request) {
	// 检查请求方法是否为POST
	if r.Method != http.MethodPost {
		// 设置允许的请求方法
		w.Header().Set("Allow", http.MethodPost)
		// 返回405 Method Not Allowed
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// 模拟数据（实际项目中应从请求体获取）
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7 // 有效期（天）

	// 将数据插入到数据库
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err) // 插入失败时返回500
		return
	}

	// 重定向到新创建的代码片段页面
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
