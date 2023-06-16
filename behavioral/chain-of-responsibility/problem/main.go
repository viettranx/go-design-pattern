package main

import "fmt"

type WebCrawler struct{}

func (WebCrawler) Crawl(url string) {
	fmt.Println("A simple process of crawling a url:", url)
	fmt.Println("- Check the url")
	fmt.Println("- Fetch url content")
	fmt.Println("- Extract information from content")
	fmt.Println("- Save information to database")
}

func main() {
	WebCrawler{}.Crawl("http://localhost/some-page")
}
