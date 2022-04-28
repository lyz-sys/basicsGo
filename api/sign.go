package api

import (
	"crypto/sha256"
	"demo/config"
	"demo/tools"
	"encoding/hex"
	"strings"
	"unicode/utf8"
)

// Sign Encrypt maps
func Sign(requestMap map[string]interface{}) {
	newDataMap := make(map[string]string)
	for k, _ := range requestMap {
		newDataMap[k] = tools.GetStringValFromMap(requestMap, k)
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
	sortSlice := tools.SortStringMap(m)

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
