package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://medium.com",
		"https://golang.org",
		"https://emininja.com",
		"https://amazon.com",
	}

	// syntax to create a channel
	c := make(chan string)

	for _, link := range links {
		//checkLink(link)
		//creating new go routine
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }

	//infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		//go checkLink(l, c)

		// writing a anonymos function
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
		//time.Sleep(2 * time.Second)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		//sending a message to channel
		c <- link
		return
	}

	fmt.Println(link, "is working!!")
	c <- link
}
