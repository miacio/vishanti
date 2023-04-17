package model

import (
	"encoding/json"
	"time"
)

// JsonTime 用于转json格式化时间
type JsonTime time.Time

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = time.Time(jsonTime).Format("2006-01-02 15:04:05")
	return json.Marshal(stamp)
}

func (jsonTime *JsonTime) UnmarshalJSON(bt []byte) error {
	var vl string
	if err := json.Unmarshal(bt, &vl); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05", vl)
	if err != nil {
		return err
	}

	jt := JsonTime(t)
	*jsonTime = jt

	return nil
}

func (jsonTime JsonTime) Value() any {
	return time.Time(jsonTime)
}

func JsonTimeNow() *JsonTime {
	t := time.Now()
	j := JsonTime(t)
	return &j
}
