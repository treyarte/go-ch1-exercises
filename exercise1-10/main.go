package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <-fmt.Sprintf("While reading body %s: %v", url, err)
		return
	}

	writeToFile("test.txt", ch, data)

	nbytes, err := io.Copy(io.Discard, resp.Body)
	
	resp.Body.Close()
	
	if err != nil {
		ch <-fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func writeToFile(fileName string, ch chan<- string, data []byte) {
	file, err := os.Create(fileName)
	if err != nil {
		ch <-fmt.Sprintf("While creating file %s %v\n", fileName, err)
	}

	defer file.Close()

	size, err:= file.Write(data);
	if err != nil {
		ch <-fmt.Sprintf("While writing to file %s %v\n", fileName, err)
	}

	fmt.Printf("Success file size:  %d bytes\n", size)
}