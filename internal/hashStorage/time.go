package hashStorage

import (
	"encoding/json"
	"strings"
	"time"
)

type aTimeDuration struct {
	duration int64
}

func (t *aTimeDuration) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	aComplex, _ := time.ParseDuration(s)
	t.duration = int64(aComplex.Seconds())
	return nil
}

func (t *aTimeDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.duration)
}
