package hashStorage

import (
	"strings"
	"time"
)

type aTime struct {
	t time.Time
}

func (t *aTime) UnmarshalJSON(data []byte) error {
	// 	format time  "2006-01-02 15:04:05"
	s := strings.Trim(string(data), "\"")
	aTime, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	t.t = aTime
	return nil
}

type aTimeDuration struct {
	duration int64
}

func (t *aTimeDuration) UnmarshalJSON(data []byte) error {
	//format time "0h1m1s"
	s := strings.Trim(string(data), "\"")
	aComplex, _ := time.ParseDuration(s)
	t.duration = int64(aComplex.Seconds())
	return nil
}
