package httpclient

import (
	"errors"
	"net/http"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

func (c *Client) RegisterUser(u *model.User) error {
	
	b, err := c.encoder(u)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/users", http.MethodPost, b)
	if err != nil {
		return err
	}
	
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return errors.New(c.error(resp.Body))
	}

	return nil
}

func (c *Client) LoginUser(u *model.User) (error) {
	b, err := c.encoder(u)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/sessions", http.MethodPost, b)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(c.error(resp.Body))
	}

	if err := c.saveCookie(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Whoami() (*model.User, error) {
	
	b, err := c.encoder(nil)
	if err != nil {
		return nil, err
	}

	req, err := c.prepareRequest("/private/whoami", http.MethodGet, b)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(c.error(resp.Body))
	}

	u := c.responseUser(resp.Body)

	return u, nil
}
