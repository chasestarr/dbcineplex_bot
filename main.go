package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

var fileName = "/cache.txt"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func wd() string {
	dir, err := os.Getwd()
	check(err)

	return dir
}

func readCache() map[string]bool {
	dir := wd()

	data, err := ioutil.ReadFile(dir + fileName)
	check(err)

	guids := strings.Split(string(data), "\n")
	m := make(map[string]bool)

	for _, id := range guids {
		m[id] = true
	}

	return m
}

func writeCache(data map[string]bool) {
	dir := wd()

	keys := []string{}
	for k := range data {
		keys = append(keys, k)
	}
	joined := strings.Join(keys, "\n")

	err := ioutil.WriteFile(dir+fileName, []byte(joined), 0644)
	check(err)
}

func itemInCache(guid string) bool {
	cache := readCache()
	_, ok := cache[guid]

	if ok {
		return true
	}

	return false
}

func addItemToCache(guid string) {
	cache := readCache()
	cache[guid] = true
	writeCache(cache)
}

func getImages(input string) []string {
	r := strings.NewReader(input)
	doc, err := html.Parse(r)
	if err != nil {
		return []string{}
	}

	var f func(*html.Node) []string
	f = func(n *html.Node) []string {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					return []string{attr.Val}
				}
			}
		}

		images := []string{}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			images = append(images, f(c)...)
		}
		return images
	}
	return f(doc)
}

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://daysbeforethecineplex.tumblr.com/rss")
	check(err)

	// cacheVals := readCache()
	// cacheVals["hello world"] = true
	// writeCache(cacheVals)

	for _, item := range feed.Items {
		// item.Title
		// item.Link
		// item.Guid
		images := getImages(item.Description)
		fmt.Println(images)
		if !itemInCache(item.GUID) {
			addItemToCache(item.GUID)
		}
	}
}
