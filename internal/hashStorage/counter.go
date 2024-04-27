package hashStorage

import "time"

type Counter struct {
	id             int
	timeSeries     [][]int
	expirationTime int
	timer          int
	stopTimer      chan bool
}

func (c *Counter) timerCounter() {
	defer close(c.stopTimer)
	for {
		select {
		case <-time.After(1 * time.Minute):
			c.timer += 1
			if c.timer > c.expirationTime {
				c.timer = 0
			}
			l := len(c.timeSeries[0])
			for i := 0; i < l; i++ {
				c.timeSeries[c.timer][i] = 0
			}
		case ok := <-c.stopTimer:
			if ok {
				break
			}
		}
	}
}

func NewCounter(timePer, expiration int) Counter {
	tmp := make([][]int, expiration)
	for i := range tmp {
		tmp[i] = make([]int, timePer)
	}
	c := Counter{timeSeries: tmp,
		expirationTime: expiration,
	}
	go c.timerCounter()
	return c
}

func (c *Counter) IncreasingTheCounterValue(value int) {
	if (c.expirationTime) >= c.timer {
		index := c.timer
		l := len(c.timeSeries[index]) - 1
		for i := 0; i < l; i++ {
			c.timeSeries[index][i] = c.timeSeries[index][i+1]
		}
		c.timeSeries[index][l] = value
	}
}

func (c *Counter) LenTimeSeries() int {
	return len(c.timeSeries)
}

func (c *Counter) SumActual() int {
	result := 0
	for _, series := range c.timeSeries {
		for _, ser := range series {
			result += ser
		}
	}
	return result
}
