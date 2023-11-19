package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(site string) bool {
	return site != "waat://furhurterwe.geds"
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	wantResults := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	gotResults := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(gotResults, wantResults) {
		t.Errorf("got %v, want %v", gotResults, wantResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "https://yandex.com"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
