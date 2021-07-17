package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const (
	normalFmt = "2006-01-02 15:04:05"
)

// 时间戳格式
type TimeStamp struct {
	time.Time
}

func (t TimeStamp) MarshalJSON() ([]byte, error) {
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (t TimeStamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *TimeStamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeStamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 时间正常格式
type TimeNormal struct {
	time.Time
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(normalFmt))
	return []byte(formatted), nil
}

func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
