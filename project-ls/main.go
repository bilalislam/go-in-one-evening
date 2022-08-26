package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var hidden = flag.Bool("a", false, "show hidden files")

func main() {
	flag.Parse()
	files := listFiles("./testdata", *hidden)
	for _, e := range files {
		fmt.Println(e)
	}
}

func listFiles(dirname string, hidden bool) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !hidden {
			if !strings.HasPrefix(f.Name(), ".") {
				dirs = append(dirs, f.Name())
			}
		} else {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}
