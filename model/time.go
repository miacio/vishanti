package model

import (
	"fmt"
	"time"
)

// JsonTime 用于转json格式化时间
type JsonTime time.Time

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
