package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Our SpyResponseWriter implements http.ResponseWriter so we can use it in the test.
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// mocking struck to imitate real store
type SpyStore struct {
	response string
	t        *testing.T
}

// implementing Store interface in our spy
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result strings.Builder
		for _, r := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store is cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result.WriteString(string(r))
			}
		}
		data <- result.String()
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		//istantiating the server with our mocked dep
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		// pretty much return an empty response that implements the writer interface
		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)

	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		// create a request with holds its own context
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		//create a derived context which olds the ability to
		//  Close the Done Chan when the func cancel is called
		cancellingCtx, cancel := context.WithCancel(request.Context())
		//after 5 mills invoce the cancel func of the new Ctx
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		// pretty much return an empty response that implements the writer interface
		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)
		if response.written {
			t.Error("a response should not have been written")
		}

	})
}
