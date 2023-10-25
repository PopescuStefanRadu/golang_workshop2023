package crawler_v1_test

import (
	"fmt"
	"sync"
	"testing"
)

type Cache struct {
	mu sync.Mutex
	m  map[string]struct{}
}

func (c *Cache) checkIfFetched(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.m[url]
	c.m[url] = struct{}{}
	return ok
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache, wg *sync.WaitGroup) {

	if depth <= 0 {
		return
	}

	if cache.checkIfFetched(url) {
		fmt.Printf("ignoring url: %s\n", url)
		return // TODO o discutie intreaga de ce mai complicat
	}
	body, urls, err := fetcher.Fetch(url) //
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher, cache, wg)
		}(u)
	}
	return
}

func TestCrawler(t *testing.T) {
	var wg sync.WaitGroup
	Crawl("https://golang.org/", 4, fetcher, &Cache{m: map[string]struct{}{}}, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
