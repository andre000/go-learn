package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "foo://bar" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"foo://bar",
		"http://yahoo.com",
	}

	expected := map[string]bool{
		"http://google.com": true,
		"foo://bar":         false,
		"http://yahoo.com":  true,
	}

	received := CheckWebsites(mockWebsiteChecker, websites)
	assertMap(t, received, expected)
}

func assertMap(t *testing.T, received map[string]bool, expected map[string]bool) {
	t.Helper()
	if !reflect.DeepEqual(received, expected) {
		t.Errorf("❌ received %v expected %v", received, expected)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "URL"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
