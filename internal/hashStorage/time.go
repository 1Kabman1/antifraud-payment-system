package hashStorage

import (
	"strings"
	"time"
)

type timeDuration struct {
	Duration int
}

//format time "0h1m1s"

func (t *timeDuration) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	aComplex, _ := time.ParseDuration(s)
	t.Duration = int(aComplex.Seconds())
	return nil
}
