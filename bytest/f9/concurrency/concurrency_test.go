package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url != "http://localhost:8000/" {
		return true
	}
	return false
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"http://localhost:8000/",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"http://localhost:8000/":     false,
		"waat://furhurterwe.geds":    true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
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
