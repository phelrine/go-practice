package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func listDir(arg string) {
	info, err := os.Stat(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if info.IsDir() {
		files, err := ioutil.ReadDir(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, info := range files {
			entry := info.Name()
			if info.IsDir() {
				entry += "/"
			}
			fmt.Println(entry)
		}
	} else {
		fmt.Println(info.Name())
	}
}

func main() {
	if len(os.Args) < 2 {
		listDir(".")
	} else {
		for _, arg := range os.Args[1:] {
			listDir(arg)
		}
	}
	os.Exit(0)
}
