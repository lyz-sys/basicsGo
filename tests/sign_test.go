package tests

import (
	"demo/api"
	"testing"
)

func TestSign(t *testing.T) {
	//参数
	requestMap := make(map[string]interface{})
	requestMap["accessToken"] = "DO0D7B4AF995199585801A939A0000"
	requestMap["requestedAppId"] = "yd-jj"
	requestMap["params"] = "{\"jtAppId\":\"dzr-gg\",\"jtAppSecret\":\"173C6F9C72FD5884CC67510D5FAOOO\"}"
	requestMap["userID"] = 298347698273984
	requestMap["mobile"] = "10011110678"
	api.Sign(requestMap)
	t.Log(requestMap)
}
