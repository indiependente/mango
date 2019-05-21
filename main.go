package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const (
	base        = "https://www.mangaeden.com"
	urlTemplate = "/it/it-manga/MANGA/CHAP/PAGE/"
)

func main() {
	manga := os.Args[1]
	manga = strings.ToLower(manga)
	manga = strings.Replace(manga, " ", "-", -1)
	url := strings.Replace(urlTemplate, "MANGA", manga, -1)
	chap := os.Args[2]
	url = strings.Replace(url, "CHAP", chap, -1)
	page := os.Args[3]
	url = strings.Replace(url, "PAGE", page, -1)

	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	pagesColl := colly.NewCollector()
	// pagesColl.Limit(&colly.LimitRule{
	// 	Parallelism: 2,
	// 	RandomDelay: 5 * time.Second,
	// })

	pagesColl.OnRequest(func(r *colly.Request) {
		fmt.Println(">", r.URL.String())
	})

	pagesColl.OnError(func(r *colly.Response, err error) {
		fmt.Println("pagesColl failed- URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	pagesColl.OnHTML("#mainImg", func(e *colly.HTMLElement) {
		fmt.Printf("%s\n", e.Attr("src"))
	})

	c.OnHTML("#pageSelect", func(e *colly.HTMLElement) {
		e.ForEach("option", func(_ int, o *colly.HTMLElement) {
			fmt.Printf("Visiting %s\n", base+o.Attr("value"))
			pagesColl.Visit(base + o.Attr("value"))
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("c failed- URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(base + url)
	c.Wait()

}
