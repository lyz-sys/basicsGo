package tools

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

// GetStringValFromMap Get the map elements as a string
func GetStringValFromMap(m map[string]interface{}, key string) string {
	val, ok := m[key]

	if !ok {
		return ""
	}

	if v, ok := val.(json.Number); ok {
		val, _ := v.Int64()
		return fmt.Sprintf("%d", val)
	}

	if v, ok := val.(int64); ok {
		return fmt.Sprintf("%d", v)
	}

	if v, ok := val.(int); ok {
		return fmt.Sprintf("%d", v)
	}

	if v, ok := val.(float64); ok {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	if v, ok := val.(float32); ok {
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	}

	if v, ok := val.(string); ok {
		return v
	}

	bs, _ := json.Marshal(val)
	return string(bs)
}

// SortStringMap Sort string maps
func SortStringMap(data map[string]string) [][]string {
	ret := make([][]string, len(data))
	rawSlice := make([]string, len(data))
	j := 0
	for k := range data {
		rawSlice[j] = k
		j++
	}
	sort.Strings(rawSlice)

	for i, k := range rawSlice {
		v, _ := data[k]
		ret[i] = []string{k, v}
	}

	return ret
}
