package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var urls []string
var processed int
var wg sync.WaitGroup

func loadUrls() {
	file, _ := os.Open("top_100.csv")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
}

func doGet(wg *sync.WaitGroup, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s ---> %s\n", url, resp.Status)
	}
	wg.Done()
}

func pool(start int) {
	finish := start + 5
	wg.Add(5)
	for _, url := range urls[start:finish] {
		go doGet(&wg, url)
	}
}
func main() {
	loadUrls()
	for i := 0; i < len(urls); i += 5 {
		pool(i)
	}
	wg.Wait()
	fmt.Println("done!")
}
