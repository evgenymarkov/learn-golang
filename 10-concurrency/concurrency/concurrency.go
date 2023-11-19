package concurrency

type result struct {
	string
	bool
}

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, site := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(site)
	}

	for range urls {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
