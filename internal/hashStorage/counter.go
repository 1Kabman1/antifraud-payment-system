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
		if ord.T.Duration > 0 && ord.T.Duration < int(time.Now().Unix()) {
			next := i.Prev()
			c.Values.Remove(i)
			i = next
		} else {
			i = i.Prev()
		}
	}
}
