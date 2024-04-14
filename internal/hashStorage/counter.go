package hashStorage

import (
	"container/list"
	"time"
)

type Counter struct {
	id         int
	TotalValue int
	Values     *list.List
}

func NewCounter() Counter {
	return Counter{}
}

func (c *Counter) DeleteExpiredOnes() {
	for i := c.Values.Back(); i != nil; {
		ord := i.Value.(Order)
		if ord.T.DurationSec > 0 && ord.T.DurationSec < int(time.Now().Unix()) {
			next := i.Prev()
			c.Values.Remove(i)
			i = next
			c.TotalValue -= ord.Value
		} else {
			i = i.Prev()
		}
	}
}
