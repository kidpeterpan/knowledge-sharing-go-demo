package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type link []string

func createLinkFromTheFile(filename string) link {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error on read csv file.\n[ERROR] -", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), "\n")
	return link(s)
}
