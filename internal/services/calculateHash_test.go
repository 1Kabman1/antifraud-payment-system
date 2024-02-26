package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	expected := map[string]string{"a": "1", "b": "1"}
	actual := calculateHash(expected)
	ok := assert.Equal(t, actual["a"], actual["b"])
	if !ok {
		t.Fatal()
	}

}
