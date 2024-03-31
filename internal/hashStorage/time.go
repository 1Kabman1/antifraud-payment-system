package hashStorage

import (
	"strings"
	"time"
)

type aTimeDuration struct {
	Duration int64
}

//format time "0h1m1s"

func (t *aTimeDuration) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	aComplex, _ := time.ParseDuration(s)
	t.Duration = int64(aComplex.Seconds())
	return nil
}
