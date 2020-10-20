package concurrency

import (
	"fmt"
	"net/http"
	"sync"
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

type (
	//Counter contain an integer counter
	Counter struct {
		mux   sync.Mutex
		value int
	}
)

//NewCounter return a new counter, started by 0
func NewCounter() *Counter {
	return &Counter{}
}

//Value return current value of counter
func (r *Counter) Value() int {
	return r.value
}

//Inc increase counter by 1
func (r *Counter) Inc() {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.value++
}
