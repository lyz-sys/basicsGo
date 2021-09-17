package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var file = "./test.json"

type user struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
	} `json:"social"`
}

func main() {
	content, _ := ioutil.ReadFile(file)
	result := struct {
		Users []user `json:"users"`
	}{}
	err := json.Unmarshal(content, &result)
	if err != nil {
		fmt.Println("Error")
	}
	for _, v := range result.Users {
		fmt.Println("name", v.Name)
		fmt.Println("type", v.Type)
		fmt.Println("age", v.Age)
		fmt.Println("social.Facebook", v.Social.Facebook)
		fmt.Println("social.Twitter", v.Social.Twitter)
		fmt.Println()
	}
}
