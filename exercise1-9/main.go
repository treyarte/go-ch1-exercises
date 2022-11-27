package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {		
		url = addHttpsToUrl(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("STATUS CODE: %d\n", resp.StatusCode)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: readings %s: STATUS CODE: %d %v\n", url, resp.StatusCode, err)
			os.Exit(1)
		}		
	}
}

func addHttpsToUrl(url string) string {
	if(strings.HasPrefix(url,"https://") || strings.HasPrefix(url,"http://")) {
		return url
	} else {
		url = "https://" + url;
		return url
	}
}