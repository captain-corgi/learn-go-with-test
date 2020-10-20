package concurrency

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://localhost:8080" {
		return true
	}
	return false
}

func TestCheckWebsites(t *testing.T) {
	type args struct {
		wc   WebsiteChecker
		urls []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "1. TDD",
			args: args{
				wc:   nil,
				urls: []string{},
			},
			want: map[string]bool{},
		},
		{
			name: "2. Map not contain valid url",
			args: args{
				wc: mockWebsiteChecker,
				urls: []string{
					"google.com",
				},
			},
			want: map[string]bool{
				"google.com": false,
			},
		},
		{
			name: "3. Map contain valid url",
			args: args{
				wc: mockWebsiteChecker,
				urls: []string{
					"google.com",
					"http://localhost:8080",
				},
			},
			want: map[string]bool{
				"google.com":            false,
				"http://localhost:8080": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWebsites(tt.args.wc, tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWebsites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestRacer(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, gotErr := Racer(slowURL, fastURL)

		if gotErr != nil {
			t.Errorf("expected no error but got %w", gotErr)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("timeout case", func(t *testing.T) {
		timeoutServer := makeDelayedServer(2 * time.Second)
		defer timeoutServer.Close()

		_, gotErr := Racer(timeoutServer.URL, timeoutServer.URL)

		if gotErr == nil {
			t.Errorf("expected an error but got none")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}

func TestCounter(t *testing.T) {
	assertCounter := func(t *testing.T, counter *Counter, want int) {
		if counter.Value() != want {
			t.Errorf("got %d, want %d", counter.Value(), want)
		}
	}
	t.Run("1. Incrementing the counter 3 times leaves it at 3 (Fixed)", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("2. Incrementing the counter n times leaves it at n (Dynamic)", func(t *testing.T) {
		wantedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCounter)
		for i := 0; i < wantedCounter; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCounter)
	})
}
