package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	var dirname string
	var level int
	flag.StringVar(&dirname, "d", "./", "dir")
	flag.IntVar(&level, "l", 0, "level")
	flag.Parse()
	listFiles(dirname, level)
}

func listFiles(dirname string, level int) {
	dirname, err := filepath.Abs(dirname)
	if err != nil {
		fmt.Println(err)
		panic("Error")
	}
	s := "|--"
	for i := 0; i < level+1; i++ {
		s = "|   " + s
	}

	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		fmt.Printf("%s%s\n", s, fi.Name())
		if fi.IsDir() {
			filename := dirname + "/" + fi.Name()
			listFiles(filename, level+1)
		}
	}
}
