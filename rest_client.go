package cas

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RestClient is an interface to use with some client to post
type RestClient interface {
	Post(string, url.Values) (*Response, error)
}

// Client is just a struct to post something
type Client struct{}

// Response is a simples struct with Status, Header and Body
type Response struct {
	Status int
	Header map[string][]string
	Body   string
}

// Post some data to a URL and
// return a pointer of our Response
// with this way we can close http client
// in the same func that's open it
func (c Client) Post(url string, params url.Values) (*Response, error) {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: transport}

	clientResponse, err := client.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	defer clientResponse.Body.Close()

	body, err := ioutil.ReadAll(clientResponse.Body)
	if err != nil {
		return nil, err
	}

	response := Response{
		Status: clientResponse.StatusCode,
		Header: clientResponse.Header,
		Body:   string(body),
	}

	return &response, nil
}
