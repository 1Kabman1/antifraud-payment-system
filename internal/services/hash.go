package services

import (
	"crypto/md5"
)

func MD5(data chan []string, ch chan<- [][16]byte) {
	var s [][16]byte
	for _, agr := range <-data {
		h := md5.Sum([]byte(agr))
		s = append(s, h)

	}
	ch <- s

}
