package models

import "time"

type Comment struct {
	Id      int
	Content string
	Created time.Time
	PostId  int
}

func (m *Comment) TableName() string {
	return TableName("comment")
}
