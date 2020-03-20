package http

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func ExampleProxyFromEnvironment() {
	r := &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "127.0.0.1:1234",
		},
	}
	nr, err := http.ProxyFromEnvironment(r)
	fmt.Println(nr, err)
	// Output:
	// <nil> <nil>
}

func TestProxyFromEnvironment(t *testing.T) {
	r := &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "127.0.0.1",
		},
	}
	nr, err := http.ProxyFromEnvironment(r)
	t.Log(nr, err)
}
