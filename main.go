package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	urlTemplate = "https://www.mangaeden.com/it/it-manga/MANGA/CHAP/1/"
)

func main() {
	manga := os.Args[1]
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, " ", "-", -1)
	url := strings.Replace(urlTemplate, "MANGA", manga, -1)
	chap := os.Args[2]
	url = strings.Replace(url, "CHAP", chap, -1)
	resp, err := http.Get(url)
	if err != nil {
		log.Panicf("Could not GET %s\n: %v\n", url, err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("Could not read Body: %v\n", err)
	}
	fmt.Println(string(data))
}
