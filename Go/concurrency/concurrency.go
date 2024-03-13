package concurrency

type WebsiteChecker func(string) bool

type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
    // creating a channel of type "result"
    resultChannel := make(chan result)

	for _, url := range urls {
        go func(u string) {
            // saving result to the channel concurrently
            // send statement
            resultChannel <- result{u, wc(u)}
        }(url)
	}

    for i := 0; i < len(urls); i++ {
        // saves to map synchronous
        // receive expression
        r := <-resultChannel
        results[r.string] = r.bool
    }

	return results
}
