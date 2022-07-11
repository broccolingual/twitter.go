package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

func downloadParallel(urls []string) {
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go downloadFromURL(u, &wg)
	}

	wg.Wait()
}

func downloadFromURL(_url string, wg *sync.WaitGroup) {
	u, err := url.Parse(_url)
	if err != nil {
		log.Fatal(err)
	}
	path := u.Path
	segments := strings.Split(path, "/")
	fileName := "./media/" + segments[len(segments)-1]
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Get(_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	size /= 1024
	defer file.Close()
	wg.Done()
}
