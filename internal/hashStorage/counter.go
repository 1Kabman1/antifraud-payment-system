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
	for i := c.Values.Front(); i != nil; i = i.Next() {
		if i.Value.(Order).t.Duration > 0 && i.Value.(Order).t.Duration < time.Now().Unix() {
			c.Values.Remove(i)
		}
	}
}
