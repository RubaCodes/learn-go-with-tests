package racer

import (
	"fmt"
	"net/http"
	"time"
)

func ConfigurableRacer(u, r string,timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(u):
		return u, nil
	case <-ping(r):
		return r ,nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", u, r)
	}
}

var tenSecondTimeout = 10 * time.Second
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}


func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// A chan struct{} is the smallest data type available 
// from a memory perspective so we get no allocation versus a bool.
// When you use var the variable will be initialised with the "zero" value of the type.
//  So for string it is "", int it is 0, etc.For channels the zero value is nil
//  and if you try and send to it with <- it will block forever because you cannot send to nil channels

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
