package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://apple.com",
		"http://microsoft.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checklink(link, c)
		//fmt.Println(link, <-c)
	}

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	//for i := 0; i < len(links); i++ {
	for l := range c { //range c is equal to <-c blocking channel to wait for
		//receiving something from channel, its basically infinite for loop for channel
		//fmt.Println(<-c)
		//go checklink(<-c, c)
		//go checklink(l, c)
		go (func(l string) {
			time.Sleep(5 * time.Second)
			checklink(l, c)
		})(l)
	}
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		//c <- "Sad it is down"
		c <- link
	} else {
		fmt.Println(link, "is up")
		//c <- "Yep it is up"
		c <- link
	}
}