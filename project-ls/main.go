package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files := listFiles("./testdata")
	for _, e := range files {
		fmt.Println(e)
	}
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
