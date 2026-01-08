// Package models 包含应用程序的数据模型和数据库操作
package models

import (
	"database/sql" // SQL数据库包
	"errors"       // 错误处理包
	"time"         // 时间处理包
)

// Snippet 结构体表示代码片段的数据模型
type Snippet struct {
	ID      int       // 代码片段的唯一标识符
	Title   string    // 代码片段的标题
	Content string    // 代码片段的内容
	Created time.Time // 创建时间
	Expires time.Time // 过期时间
}

// SnippetsModel 结构体用于与snippets表进行数据库交互
type SnippetsModel struct {
	DB *sql.DB // 数据库连接池
}

// Insert 方法用于向数据库中插入新的代码片段
// 参数：
//   - title: 代码片段的标题
//   - content: 代码片段的内容
//   - expires: 过期时间（天）
// 返回：
//   - int: 新插入记录的ID
//   - error: 可能的错误
func (m *SnippetsModel) Insert(title, content string, expires int) (int, error) {
	// SQL插入语句
	// 使用UTC_TIMESTAMP()获取当前UTC时间
	// 使用DATE_ADD()计算过期时间
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// 执行SQL语句
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// 获取插入记录的ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 将int64类型的ID转换为int并返回
	return int(id), nil
}

// Get 方法用于根据ID获取代码片段
// 参数：
//   - id: 代码片段的ID
// 返回：
//   - *Snippet: 代码片段指针
//   - error: 可能的错误
func (m *SnippetsModel) Get(id int) (*Snippet, error) {
	// SQL查询语句
	// 只查询未过期的代码片段（expires > UTC_TIMESTAMP()）
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// 执行查询，获取单行结果
	row := m.DB.QueryRow(stmt, id)

	// 创建Snippet实例
	s := &Snippet{}

	// 将查询结果扫描到Snippet实例中
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// 检查是否是记录未找到的错误
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// 返回获取到的代码片段
	return s, nil
}

// Latest 方法用于获取最近创建的10个未过期代码片段
// 返回：
//   - []*Snippet: 代码片段指针切片
//   - error: 可能的错误
func (m *SnippetsModel) Latest() ([]*Snippet, error) {
	// SQL查询语句
	// 查询未过期的代码片段，按ID降序排序，限制10条
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// 执行查询，获取多行结果
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	
	// 延迟关闭结果集，确保资源释放
	defer rows.Close()

	// 创建代码片段切片
	snippets := []*Snippet{}

	// 遍历结果集
	for rows.Next() {
		// 创建Snippet实例
		s := &Snippet{}
		
		// 将查询结果扫描到Snippet实例中
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		
		// 将Snippet实例添加到切片中
		snippets = append(snippets, s)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// 返回获取到的代码片段切片
	return snippets, nil
}
