package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

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

const (
	defaultTimeout = 1 * time.Second
)

//Racer return the fastest url within default timeout
func Racer(a string, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

//ConfigurableRacer return the fastest url within input timeout
func ConfigurableRacer(a string, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
