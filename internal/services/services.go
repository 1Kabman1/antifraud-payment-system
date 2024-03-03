package services

import (
	"crypto/md5"
	"errors"
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"strconv"
	"strings"
)

// calculateHash - hash function
func calculateHash(data map[int]string) map[int][16]byte {
	result := make(map[int][16]byte)

	for key, val := range data {
		h := md5.Sum([]byte(val))
		result[key] = h
	}

	return result

}

// prepareTheDataForHashing - prepares data for hashing
func prepareTheDataForHashing(rules map[int]*hashStorage.Rule, payment map[string]interface{}) (map[int]string, error) {
	aggregatesBy := make(map[int]string, len(rules))
	var aBuilder strings.Builder

	for _, rule := range rules {
		aggregate := ""
		var flag bool

		for _, agg := range rule.AggregateBy {
			v, ok := payment[agg]
			if !ok {
				flag = true
				aBuilder.Reset()
				break
			}

			switch aInterface := v.(type) {
			case float64:
				intInterface := int(aInterface)
				_, err := aBuilder.WriteString(strconv.Itoa(intInterface))
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

		}

		if !flag {
			aBuilder.WriteString(strconv.Itoa(rule.AggregationRuleId))
			aggregate = aBuilder.String()
			aggregatesBy[rule.AggregationRuleId] = aggregate
			aBuilder.Reset()
		}

	}
	return aggregatesBy, nil
}
