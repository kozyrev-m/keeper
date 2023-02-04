package httpclient

import (
	"errors"
	"log"
	"net/http"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

// RegisterUser creates new user.
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

	log.Printf("successfully created user with login '%s' in system keeper!\n", u.Login)

	return nil
}

// LoginUser creates session user.
func (c *Client) LoginUser(u *model.User) error {
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

	log.Printf("successfully created session for user: '%s'!\n", u.Login)

	return nil
}

// Whoami gets user data.
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
