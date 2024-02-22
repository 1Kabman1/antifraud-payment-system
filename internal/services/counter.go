package services

type counter struct { // пока оставить
	//id     int // не обязательно
	//count  int //
	//amount int // value always int
	value int
}

func newCounter() counter {
	return counter{}
}
