package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var statuses = make(map[string]int)
var wg sync.WaitGroup

func doGet(url string) {
	timeStart := time.Now()
	resp, _ := http.Get(url)

	fmt.Println(url, resp.Status, time.Since(timeStart))
	statuses[resp.Status] += 1
	wg.Done()
}

func readURLS(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

func main() {

	urls, err := readURLS("urls.txt")
	if err != nil {
		fmt.Printf("readURLS: %s\n", err)
	}
	for _, url := range urls {
		wg.Add(1)
		go doGet(url)
	}

	wg.Wait()

	fmt.Println("Done! (╯°□°)╯︵ ┻━┻")
	for status, count := range statuses {
		fmt.Printf("Total urls with status %s: %d\n", status, count)
	}

}
