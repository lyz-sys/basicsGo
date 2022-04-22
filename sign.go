package main

import (
	"crypto/sha256"
	"demo/config"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	//参数
	requestMap := make(map[string]interface{})
	requestMap["accessToken"] = "DO0D7B4AF995199585801A939A0000"
	requestMap["requestedAppId"] = "yd-jj"
	requestMap["params"] = "{\"jtAppId\":\"dzr-gg\",\"jtAppSecret\":\"173C6F9C72FD5884CC67510D5FAOOO\"}"
	requestMap["userID"] = "298347698273984"
	requestMap["mobile"] = "10011110678"
	sign(requestMap)
	fmt.Println(requestMap)
}

func sign(requestMap map[string]interface{}) {
	newDataMap := make(map[string]string)
	for k, _ := range requestMap {
		newDataMap[k] = getStringValFromMap(requestMap, k)
	}
	keyStr := map2Str(newDataMap)

	// 去除双引号、全角空格
	keyStr += strings.Replace(config.AppKey, "\"", "", -1)
	keyStr = strings.Replace(keyStr, "　", " ", -1)
	// 加密
	sha256New := sha256.New()
	sha256New.Write([]byte(keyStr))
	sha256String := hex.EncodeToString(sha256New.Sum(nil))
	requestMap["sign"] = sha256String
	return
}

func map2Str(m map[string]string) (str string) {
	sortSlice := sortStringMap(m)

	for _, k := range sortSlice {
		str += k[0]
		str += "="
		str += k[1]
		str += "&"
	}
	r, size := utf8.DecodeLastRuneInString(str)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	str = str[:len(str)-size]
	return
}

func sortStringMap(data map[string]string) [][]string {
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

func getStringValFromMap(m map[string]interface{}, key string) string {
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
