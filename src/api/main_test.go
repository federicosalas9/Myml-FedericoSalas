package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func BenchmarkApiGo(b *testing.B) {

	router := Start()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/myml/1", nil)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for n := 0; n < 100; n++ {
		go func() {
			router.ServeHTTP(w, req)
			wg.Done()
		}()
	}
	wg.Wait()
}
