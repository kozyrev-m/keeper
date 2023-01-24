package httpclient

import (
	"net/http"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

func (c *Client) RegisterUser(u *model.User) (error) {
	
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

	if err := c.saveCookie(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Whoami() (error) {
	
	b, err := c.encoder(nil)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/whoami", http.MethodGet, b)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// payload, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	return nil
}