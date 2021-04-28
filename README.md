# http-moq

*Interfacing All The Things* is [bad practice](https://about.sourcegraph.com/go/idiomatic-go/#Interface-All-The-Things), 
and yet we have several packages that are dedicated to allowing you to create and mock interfaces in Go 
(including [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter), which is used in this package). 
This is often used where it shouldn't be - mocking your HTTP client wrappers with something like the following:
```go
type Client interface {
  GetSomething(...)
  CreateSomething(...)
  UpdateSomething(...)
  DeleteSomething(...)
}
```
Instead, use HTTP Moq to mock out the underlying *HTTP client* that is used by your Client package. Get rid of all your
interfaces! Get rid of all your fakes! Rely only on a mocked response to test higher level implementation.

So, you have a package `foo` which contains a client that is already being correctly tested
against a mock server. In it you have the following client implementation:
```go
package foo

func NewClient() *Client {
  return &Client{
    url: "http://foo.com",
    c:   http.DefaultClient,
  }
}

type Client struct {
  url string
  c   *http.Client
}

type Something struct {}

// GetSomething makes an HTTP call to the Foo service and returns
// an instance of Something.
func (c *Client) GetSomething(id string) (Something, error) {
  res, err := c.c.Get(c.url)  
  if err != nil {
    // Handle error.
  }
  // Actual HTTP implementation goes here.
  return Something{}, nil
}

// WithClient sets the underlying HTTP client for Foo to use.
func (c *Client) WithClient(client *http.Client) {
  c.c = client
}
```
No interfaces, no mocking, no fakes.

Now, in order to test in some higher level package using HTTP Moq, all you have to do is something like this.
```go
// Generate a mock HTTP client.
fakeHTTPClient := httpmoq.NewClient()
// Generate a real Foo client.
fooClient := foo.NewClient()
// Set Foo's underlying HTTP client to be the mocked client.
fooClient.WithClient(fakeHTTPClient)
```
Now, when you test you can use familiar methods found in [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter)
to mock errors and responses.

It's a bit lower-level, but it *mostly* gets you away from the Go Anti-pattern of Interfacing everything,
even if this package itself is creating an interface where it shouldn't ;).
