package base

import (
	"encoding/base64"
	"encoding/json"
)

type (
	// 传统的分页对象
	Pager struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	// Cursor方式分页码对象
	Next string
	// Cursor分页参数对象
	Cursor struct {
		// 是否有效
		Valid bool `json:"-"`
		// 指定条件字段，其字段类型必须为可排序的。例："id"， "user.id"，"user.login_time"
		Field string `json:"field"`
		// 该值为排序字段的具体值。例："123"，"2023-07-27 15:41:46"，"abc"
		CursorID any `json:"cursor_id"`
		// 分页条数，默认从config/system.json配置文件读取设置的pagesize
		Size int `json:"size"`
		// 指定排序方式。例如： "DESC"，"ASC"
		SortType string `json:"sort_type"`
	}
)

// 判断是否为空
func (c *Cursor) IsEmpty() bool {
	// 空指针
	if c == nil {
		return true
	}
	// 空结构体
	if *c == (Cursor{}) {
		return true
	}
	// 逻辑空
	if !c.Valid {
		return true
	}
	// 没值
	if c.Field == "" && c.Size == 0 && c.SortType == "" {
		return true
	}
	return false
}

// 编码为base64串的Next
func (c *Cursor) Encode() Next {
	b, err := json.Marshal(c)
	if err != nil {
		return Next("")
	}
	return Next(base64.StdEncoding.EncodeToString(b))
}

// 解码为Cursor对象
func (n Next) Decode() Cursor {
	cursor := &Cursor{}
	b, err := base64.StdEncoding.DecodeString(string(n))
	if err != nil {
		return Cursor{}
	}
	err = json.Unmarshal(b, cursor)
	if err != nil {
		return Cursor{}
	}
	cursor.Valid = true
	return *cursor
}
