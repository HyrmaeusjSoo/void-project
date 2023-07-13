package base

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// 格式化好的时间类型
type Time sql.NullTime

func NewTime(t time.Time) *Time {
	return &Time{
		Valid: true,
		Time:  t,
	}
}

func (n *Time) Scan(value interface{}) error {
	return (*sql.NullTime)(n).Scan(value)
}

func (n Time) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func (n Time) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return []byte(fmt.Sprintf("\"%v\"", n.Time.Format(time.DateTime))), nil
	}
	return json.Marshal(nil)
}

func (n *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		return nil
	}

	var timestr string
	err := json.Unmarshal(b, &timestr)
	if err != nil {
		n.Valid = false
		return nil
	}
	n.Time, err = time.ParseInLocation(time.DateTime, timestr, time.Local)
	if err == nil {
		n.Valid = true
	}
	return err
}
