package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: map[string]error{}}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if _, ok := fetched.m[url]; ok || depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)

	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)
	wg := &sync.WaitGroup{}
	for _, u := range urls {
		wg.Add(1)
		go func(wg *sync.WaitGroup, url string) {
			Crawl(url, depth-1, fetcher)
			wg.Done()
		}(wg, u)
	}
	wg.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
	//time.Sleep(time.Second * 10)
	for k, v := range fetched.m {
		if v == nil {
			fmt.Printf("existing url: %v\n", k)
		}
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
