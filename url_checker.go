package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var urls []string

func doGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s ---> %s\n", url, resp.Status)
	}
}

func loadUrls() {
	file, _ := os.Open("top_100.csv")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
}

func main() {
	loadUrls()
	for _, url := range urls {
		doGet(url)
	}
	fmt.Println("Done! (╯°□°)╯︵ ┻━┻")
}
