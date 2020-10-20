package concurrency

type (
	//WebsiteChecker is a function to check website
	WebsiteChecker func(string) bool
	result         struct {
		string
		bool
	}
)

//CheckWebsites check a list urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	rsChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			rsChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-rsChannel
		results[result.string] = result.bool
	}
	return results
}
