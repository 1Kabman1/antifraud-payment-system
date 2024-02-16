package services

import (
	"crypto/md5"
)

func MD5Slice(data []string) [16]byte {
	temp := ""
	for _, agr := range data {
		temp += agr
	}
	h := md5.Sum([]byte(temp))
	return h
}

func MD5String(data string) [16]byte {

	h := md5.Sum([]byte(data))
	return h
}
