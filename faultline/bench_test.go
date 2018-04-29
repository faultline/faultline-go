package faultline_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faultline/faultline-go/faultline"
)

func BenchmarkSendNotice(b *testing.B) {
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"data":{"errors":{"postCount":27}}}`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))

	notifier := faultline.NewNotifier("sample-project", "123", server.URL, []interface{}{})
	notifier.SetEndpoint(server.URL)

	notice := notifier.Notice(errors.New("benchmark"), nil, 0)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			count, err := notifier.SendNotice(notice)
			if err != nil {
				b.Fatal(err)
			}
			if count != 27 {
				b.Fatalf("got %q, wanted 27", count)
			}
		}
	})
}
