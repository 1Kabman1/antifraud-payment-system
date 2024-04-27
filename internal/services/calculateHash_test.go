package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	expected := map[int]string{1: "1", 2: "1"}
	actual := calculateHash(expected)
	ok := assert.Equal(t, actual[1], actual[2])
	if !ok {
		t.Fatal()
	}
}
