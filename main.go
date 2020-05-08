package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	testfile := os.Args[1]
	links := createLinkFromTheFile(testfile)
	// add
	totals := len(links)
	success := 0
	startUnixNano := time.Now().UnixNano()

	c := make(chan int)

	for _, link := range links {
		go healthcheck(link, c)
	}

	// modify
	for i := 0; i < len(links); i++ {
		success += <-c
	}
	// add
	endUnixNano := time.Now().UnixNano()
	totalTimeUnixNano := endUnixNano - startUnixNano
	failure := totals - success
	totalTimeMs := totalTimeUnixNano / int64(time.Millisecond)
	writeResult(totals, success, failure, totalTimeMs)

}

func healthcheck(link string, c chan int) { // change to int
	_, err := http.Get(link)
	// add
	success := 1
	failure := 0
	if err != nil {
		downText := link + " Might Be Down"
		fmt.Println(downText)
		c <- failure
		return
	}
	upText := link + " is up!"
	fmt.Println(upText)
	c <- success
}

// add
func writeResult(totals int, success int, failure int, totalTimeMs int64) {
	fmt.Println("Perform website checking...")
	fmt.Println("Done!")
	fmt.Println("Checked webistes: ", totals)
	fmt.Println("Successful websites: ", success)
	fmt.Println("Failure websites: ", failure)
	fmt.Println("Total times to finished checking website: ", totalTimeMs, "millisecond(s)")
}
