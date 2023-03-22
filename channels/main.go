package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	impWebsites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	//channel creation
	c := make(chan string)

	for _, website := range impWebsites {
		go checkLink(website, c)
	}

	for l := range c {
		go checkLink(l, c)
	}

	//fmt.Println(<-c)
}
func checkLink(website string, c chan string) {
	_, err := http.Get(website)
	fmt.Printf("Checking site:%s is Up!\n", website)
	c <- website
	if err != nil {
		fmt.Print("Error occured while getting the site status", err)
		c <- website
	}
	time.Sleep(6000)
}

// concurrency - making use of CPU core efficiently with multiple Go routines
//parallelism - exeuting two Go routines in parallel with more than one CPU Core
