package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetsModel struct {
	DB *sql.DB
}

func (m *SnippetsModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetsModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetsModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
