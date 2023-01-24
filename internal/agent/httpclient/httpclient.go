// Package httpclient provides http-client to send private information to server.
package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kozyrev-m/keeper/internal/agent/config"

	"github.com/juju/persistent-cookiejar"
)

// Client contains http-client implementation.
type Client struct {
	http.Client
	cookiejar *cookiejar.Jar
}

// New creates instance http-client.
func New() *Client {
	client := &Client{}

	client.initJar()

	return client
}

// initJar inits cookie jar.
func (c *Client) initJar() {
	// TODO: Move to config filename: "/tmp/keeper_cookie"
	jar, err := cookiejar.New(&cookiejar.Options{Filename: "/tmp/keeper_cookie"})
	if err != nil {
		log.Println(err)
	}

	c.Jar = jar
	c.cookiejar = jar
}

// saveCookie saves cookie to file.
func (c *Client) saveCookie() error {
	return c.cookiejar.Save()
}

// prepareRequest prepares request.
func (c *Client) prepareRequest(url string, method string, buf *bytes.Buffer) (*http.Request ,error) {
	// TODO: For "http://%s%s" to exclude an error
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("http://%s%s", config.Address, url),
		buf,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// encoder helps to encode some data to json. 
func (c *Client) encoder(data interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf, nil
}