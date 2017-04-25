package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var urls []string

func loadUrls() {
	file, _ := os.Open("top_100.csv")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
}

func doGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Sprintf("%s ---> %s", url, resp.Status)
	}
}

func main() {
	loadUrls()
	for _, url := range urls {
		go doGet(url)
	}
	fmt.Println("Done! ‎(ﾉಥ益ಥ）ﾉ﻿ ┻━┻")
}
