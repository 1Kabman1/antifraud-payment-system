package hashStorage

import (
	"time"
)

type Counter struct {
	id               int
	timeSeriesValues [][]int
	expirationTime   int
	timer            int
}

func (c *Counter) timerCounter() {
	for {
		select {
		case <-time.After(1 * time.Minute):
			c.timer += 1
			if c.timer >= c.expirationTime {
				c.timer = 0
			}
			for i := 0; i < len(c.timeSeriesValues[0]); i++ {
				c.timeSeriesValues[c.timer][i] = 0
			}
		}
	}
}

func NewCounter(timePer, expiration int) Counter {
	tmp := make([][]int, expiration)
	for i := range tmp {
		tmp[i] = make([]int, timePer)
	}
	c := Counter{timeSeriesValues: tmp,
		expirationTime: expiration,
	}
	return c
}

func (c *Counter) IncreasingTheCounterValue(value int) {
	index := c.timer
	l := len(c.timeSeriesValues[index]) - 1
	for i := 0; i < l; i++ {
		c.timeSeriesValues[index][i] = c.timeSeriesValues[index][i+1]
	}
	c.timeSeriesValues[index][l] = value
}

func (c *Counter) LenTimeSeries() int {
	return len(c.timeSeriesValues)
}

func (c *Counter) SumActual() int {
	result := 0
	for _, series := range c.timeSeriesValues {
		for _, ser := range series {
			result += ser
		}
	}
	return result
}
