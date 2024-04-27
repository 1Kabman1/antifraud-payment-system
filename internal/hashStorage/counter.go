package hashStorage

type Counter struct {
	id         int
	TotalValue int
	//	Values     *list.List // delete
	timeSeries [][]int // new
}

func NewCounter(timePeriod, expirationTime int) Counter {
	tmp := make([][]int, expirationTime)
	for i := range tmp {
		tmp[i] = make([]int, timePeriod)
	}
	c := Counter{timeSeries: tmp}
	return c
}

func (c *Counter) IncreasingTheCounterCount() {

	c.TotalValue += 1
	ord.Value = 1

}

func (c *Counter) IncreasingTheCounterAmount(amount int) {
	c.TotalValue += amount
	ord.Value = amount

}

//func (c *Counter) DeleteExpiredOnes() {
//	for i := c.Values.Back(); i != nil; {
//		ord := i.Value.(Order)
//		if ord.T.DurationSec > 0 && ord.T.DurationSec < int(time.Now().Unix()) {
//			next := i.Prev()
//			c.Values.Remove(i)
//			i = next
//			c.TotalValue -= ord.Value
//		} else {
//			i = i.Prev()
//		}
//	}
//}
