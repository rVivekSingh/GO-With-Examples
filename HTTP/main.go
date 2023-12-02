package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com/search?q=test")

	if err != nil {
		fmt.Println("Error:-", err)
		os.Exit(1)
	}

	bs := make([]byte, 1999)
	resp.Body.Read(bs)
	//fmt.Println(resp.Body.Read())
	fmt.Println(string(bs))
}
