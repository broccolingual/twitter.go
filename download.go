package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func downloadFromURL(_url string) {
	u, err := url.Parse(_url)
	if err != nil {
		log.Fatal(err)
	}
	path := u.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
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
	defer file.Close()
	fmt.Printf("Downloaded a file %s with size %d byte.", fileName, size)
}
