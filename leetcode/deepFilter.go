// https://leetcode.com/discuss/interview-question/328553/amazon-phone-screen-deep-filter
/*
Given an object and a filter function,
write a function that will go through and filter the object,
then return a filtered object
*/
package leetcode

import (
	"encoding/json"
	"fmt"
)

func deepFilter(input map[string]interface{}, callback func(element interface{}) bool) map[string]interface{} {
	filteredResult := make(map[string]interface{})

	for key, value := range input {
		if v, ok := value.(map[string]interface{}); ok {
			nestedObject := deepFilter(v, callback)
			if len(nestedObject) > 0 {
				filteredResult[key] = nestedObject
			}
		} else if callback(value) {
			filteredResult[key] = value
		}
	}

	return filteredResult
}

func DoDeepFilter() {
	intCase := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"c": 2,
			"d": -3,
			"e": map[string]interface{}{
				"f": map[string]interface{}{
					"g": -4,
				},
			},
			"h": map[string]interface{}{
				"i": 5,
				"j": 6,
			},
		},
	}
	intAns := deepFilter(intCase, func(element interface{}) bool {
		switch v := element.(type) {
		case int:
			return v >= 0
		default:
			return false
		}
	})

	stringCase := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"c": "Hello World",
			"d": 2,
			"e": map[string]interface{}{
				"f": map[string]interface{}{
					"g": -4,
				},
			},
			"h": "Good Night Moon",
		},
	}
	stringAns := deepFilter(stringCase, func(element interface{}) bool {
		switch element.(type) {
		case string:
			return true
		default:
			return false
		}
	})

	answers := []map[string]interface{}{intAns, stringAns}
	for _, answer := range answers {
		jsonString, _ := json.Marshal(answer)
		fmt.Println(string(jsonString))
	}
}
