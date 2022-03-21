package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileInfo, err := os.Stat("./test.json")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Error())
	}
	res := filepath.IsAbs("./test.json")
	abs, err := filepath.Abs("./test.json")
	fmt.Println(res, abs)
	fmt.Println(path.Join(abs, ".."))
	fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
}
