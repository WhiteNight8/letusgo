// Package main 是应用程序的入口包
package main

import (
	"database/sql" // 导入数据库SQL包
	"flag"         // 导入命令行参数解析包
	"html/template"
	"log"      // 导入日志包
	"net/http" // 导入HTTP服务器包
	"os"       // 导入操作系统包

	"letgo.snippetbox/internal/models" // 导入内部数据模型包

	_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动，但不直接使用
)

// application 结构体定义了应用程序的核心依赖
type application struct {
	infoLog       *log.Logger                   // 信息日志记录器
	errorLog      *log.Logger                   // 错误日志记录器
	snippets      *models.SnippetsModel         // 代码片段数据模型
	templateCache map[string]*template.Template // 模板缓存
}

// main 函数是应用程序的入口点
func main() {
	// 定义HTTP服务器端口命令行参数
	addr := flag.String("addr", ":4000", "HTTP网络地址")
	// 定义MySQL数据源名称命令行参数
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=True", "MySQL数据源名称")

	// 解析命令行参数
	flag.Parse()

	// 创建信息日志记录器
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// 创建错误日志记录器，包含文件名和行号
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 打开数据库连接
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err) // 数据库连接失败时退出程序
	}

	// 延迟关闭数据库连接
	defer db.Close()

	// 创建模板缓存
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err) // 模板缓存创建失败时退出程序
	}

	// 初始化应用程序结构体
	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		snippets:      &models.SnippetsModel{DB: db},
		templateCache: templateCache,
	}

	// 配置HTTP服务器
	srv := &http.Server{
		Addr:     *addr,        // 服务器监听地址
		ErrorLog: errorLog,     // 服务器错误日志
		Handler:  app.routes(), // 请求处理器
	}

	// 记录服务器启动信息
	infoLog.Printf("服务器正在端口 %s 上运行", *addr)
	// 启动HTTP服务器
	err = srv.ListenAndServe()
	errorLog.Fatal(err) // 服务器启动失败时退出程序
}

// openDB 函数用于创建并测试数据库连接
func openDB(dsn string) (*sql.DB, error) {
	// 创建数据库连接池
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 测试数据库连接是否可用
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil // 返回数据库连接池
}
