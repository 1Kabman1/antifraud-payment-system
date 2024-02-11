package services

import (
	"crypto/md5"
)

func MD5(data []string, ch chan<- [16]byte) {
	temp := ""
	for _, agr := range data {
		temp += agr
	}
	h := md5.Sum([]byte(temp))
	ch <- h
}
