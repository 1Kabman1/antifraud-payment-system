package services

import (
	"crypto/md5"
	"errors"
	"strconv"
	"strings"
)

// calculateHash - hash function
func calculateHash(data map[string]string) map[string][16]byte {
	result := make(map[string][16]byte)

	for key, val := range data {
		h := md5.Sum([]byte(val))
		result[key] = h
	}

	return result

}

// prepareTheDataForHashing - prepares data for hashing
func prepareTheDataForHashing(h map[string]interface{}, mapPING map[string]interface{}) (map[string]string, error) {
	aggregatesBy := make(map[string]string, len(h))
	var aBuilder strings.Builder

	for _, tempRule := range h {

		aRule := tempRule.(rule)
		aggregate := ""
		var flag bool

		for _, agg := range aRule.AggregateBy {
			if v, ok := mapPING[agg]; ok {
				switch aInterface := v.(type) {
				case float64:
					_, err := aBuilder.WriteString(strconv.FormatFloat(aInterface, 'E', -1, 64))
					if err != nil {
						return nil, errors.New("the type is not float")
					}
					aBuilder.WriteString(agg) // добавляю имя агрегирующего чтобы убрать совпадение
				case string:
					_, err := aBuilder.WriteString(aInterface)
					if err != nil {
						return nil, errors.New("the type is not string")
					}
					aBuilder.WriteString(agg) // добавляю имя агрегирующего чтобы убрать совпадение
				}

			} else { // Если хотябы одно агрегируемое из правила не совпадает с поступившим на почет поручением, то оно автоматом исключается
				flag = false
				aBuilder.Reset()
				break
			}

		}

		if flag {
			aggregate = aBuilder.String()
			aggregatesBy[aRule.Name] = aggregate
			aBuilder.Reset()
		}

	}
	return aggregatesBy, nil
}
