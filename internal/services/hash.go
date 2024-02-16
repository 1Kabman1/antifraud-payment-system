package services

import (
	"crypto/md5"
)

func MD5(data chan map[string]string, ch chan<- map[string][16]byte) {
	result := make(map[string][16]byte)

	for key, val := range <-data {
		h := md5.Sum([]byte(val))
		result[key] = h
	}

	ch <- result

}
