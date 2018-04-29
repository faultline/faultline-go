package faultline_test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	. "github.com/onsi/ginkgo"

	"github.com/faultline/faultline-go/faultline"
)

var _ = Describe("Notifier", func() {
	var notifier *faultline.Notifier

	BeforeEach(func() {
		handler := func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"data":{"errors":{"postCount":27}}}`))
		}
		server := httptest.NewServer(http.HandlerFunc(handler))

		notifier = faultline.NewNotifier("sample-project", "123", "https://api.example.com/v0", []interface{}{})
		notifier.SetEndpoint(server.URL)
	})

	It("is race free", func() {
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				notifier.Notify("hello", nil)
			}()
		}
		wg.Wait()

		notifier.Flush()
	})
})
