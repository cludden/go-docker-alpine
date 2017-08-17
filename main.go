package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("URL")
	length, err := pageLength(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Page length of %s: %d\n", url, length)
	os.Exit(0)
}

func pageLength(url string) (length int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	length = len(body)
	return
}
