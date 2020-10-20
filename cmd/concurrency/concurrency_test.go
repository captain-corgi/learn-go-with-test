package concurrency

import (
	"reflect"
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
