// HTTP client interface to be mocked and used for testing purposes.
//
// See https://golang.org/src/net/http/client.go for implementation.

package httpmoq

import (
	"io"
	"net/http"
	"net/url"
)

// To generate install `counterfeiter`.
//
// https://github.com/maxbrunsfeld/counterfeiter#installing-counterfeiter-to-gopathbin
//
// Then run
//
// $ counterfeiter -o . --fake-name Client internal Client
type Client interface {
	Get(string) (*http.Response, error)
	Do(*http.Request) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)
	Head(string) (*http.Response, error)
	CloseIdleConnections()
}
