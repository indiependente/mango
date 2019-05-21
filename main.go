package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
		DomainGlob:  "*mangaeden.*",
		Parallelism: 5,
	})
	pagesColl := colly.NewCollector(
		colly.Async(true),
	)
	pagesColl.Limit(&colly.LimitRule{
		DomainGlob:  "*mangaeden.*",
		Parallelism: 5,
	})

	pagesColl.OnError(func(r *colly.Response, err error) {
		fmt.Println("pagesColl failed- URL:", r.Request.URL, "\nError:", err)
	})

	pagesColl.OnHTML("#mainImg", func(e *colly.HTMLElement) {
		altFields := strings.Fields(e.Attr("alt"))
		pageNum := altFields[len(altFields)-1]

		resp, err := http.Get("https:" + e.Attr("src"))
		if err != nil {
			fmt.Printf("Failed to get %s: %v\n", "https"+e.Attr("src"), err)
			return
		}
		defer resp.Body.Close()
		f, err := os.Create(manga + "_" + chap + "_" + pageNum + ".png")
		if err != nil {
			fmt.Printf("Failed to create file: %v\n", err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			fmt.Printf("Could not copy response content: %v\n", err)
			return
		}

	})

	c.OnHTML("#pageSelect", func(e *colly.HTMLElement) {
		e.ForEach("option", func(_ int, o *colly.HTMLElement) {
			fmt.Printf("Visiting %s\n", base+o.Attr("value"))
			pagesColl.Visit(base + o.Attr("value"))
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("c failed- URL:", r.Request.URL, "\nError:", err)
	})

	c.Visit(base + url)
	c.Wait()
	pagesColl.Wait()
}
