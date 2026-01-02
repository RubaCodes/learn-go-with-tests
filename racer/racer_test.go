package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(30 * time.Millisecond)
		fastServer := makeDelayedServer(0)
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got ,err := Racer(slowServer.URL, want)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("return a error if the server doesn't respond in less than 10 seconds", func(t *testing.T) {
		serverB := makeDelayedServer(12 * time.Second)
		defer serverB.Close()
		_, err := ConfigurableRacer(serverB.URL,serverB.URL, 10 * time.Second)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))

}

