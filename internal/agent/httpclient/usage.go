package httpclient

import (
	"encoding/json"
	"io"
	"log"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

func (c *Client) error(body io.ReadCloser) string {
	res := &responseError{}
	if err := json.NewDecoder(body).Decode(res); err != nil {
		log.Printf("can't decode data from json: %s", err.Error())
	}

	return res.Error
}

func (c *Client) responseUser(body io.ReadCloser) *model.User {
	u := &model.User{}
	if err := json.NewDecoder(body).Decode(u); err != nil {
		log.Printf("can't decode data from json: %s", err.Error())
	}

	return u
}