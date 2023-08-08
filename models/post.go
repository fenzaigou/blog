package models

import "time"

type Post struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Updated time.Time
	Views   int
	IsTop   int8
	Url     string
}

func (m *Post) TableName() string {
	return TableName("post")
}
