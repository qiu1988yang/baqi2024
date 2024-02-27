package utils

import (
	"fmt"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

// ConvertArrayToMap 全局数组转换
func ConvertArrayToMap(arr interface{}, keyField string) (map[string]Person, error) {
	result := make(map[string]Person)

	for _, elem := range arr.([]Person) {
		switch keyField {
		case "ID":
			result[fmt.Sprintf("%d", elem.ID)] = elem
		case "Name":
			result[elem.Name] = elem
		case "Age":
			result[fmt.Sprintf("%d", elem.Age)] = elem
		}
	}

	return result, nil
}
