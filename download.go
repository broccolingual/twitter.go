package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"
)

func downloadParallel(urls []string) {
	var wg sync.WaitGroup
	var s = semaphore.NewWeighted(50)

	for _, u := range urls {
		wg.Add(1)
		go downloadFromURL(u, &wg, s)
	}

	wg.Wait()
}

func downloadFromURL(_url string, wg *sync.WaitGroup, s *semaphore.Weighted) {
	defer wg.Done()
	if err := s.Acquire(context.Background(), 1); err != nil {
		return
	}
	defer s.Release(1)

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
	fmt.Printf("%s %dKB\n", fileName, size/1024)
	defer file.Close()
}
