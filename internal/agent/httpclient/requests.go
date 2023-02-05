package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/kozyrev-m/keeper/internal/agent/httpclient/sheme"
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

// DownloadFile gets file from server by name.
func (c *Client) UploadFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func () {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	} ()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", path.Base(filepath))
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/file", http.MethodPost, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// DownloadFile gets file from server by name.
func (c *Client) DownloadFile(filename string) error {

	b, err := c.encoder(nil)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest(fmt.Sprintf("/private/file/%s", filename), http.MethodGet, b)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Create the file
	filepath := fmt.Sprintf("/tmp/%s", filename)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// ListFiles gets file list from server by owner.
func (c *Client) ListFiles() error {

	b, err := c.encoder(nil)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/file", http.MethodGet, b)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	input := string(body)
	input = input[1 : len(input)-2]
	input = fmt.Sprintf("\"%s\"", input)

	jsonInput, err := strconv.Unquote(input)
	if err != nil {
		return err
	}

	respFiles := &sheme.ResponseFiles{}
	if err := json.Unmarshal([]byte(jsonInput), respFiles); err != nil {
		return err
	}

	fmt.Printf("Your (%s) file list:\n", c.User.Login)
	for _, file := range respFiles.Files {
		fmt.Printf("- %s\n", file.Name)
	}

	return nil
}

func (c *Client) AddBankCardData(bc *model.BankCard) error {
	if err := bc.Validate(); err != nil {
		return err
	}

	b, err := c.encoder(bc)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/bankcard", http.MethodPost, b)
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

	log.Println("Add bank card data!")
	return nil
}

func (c *Client) GetBankCards() error {
	b, err := c.encoder(nil)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/bankcard", http.MethodGet, b)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	input := string(body)
	input = input[1 : len(input)-2]
	input = fmt.Sprintf("\"%s\"", input)
	
	jsonInput, err := strconv.Unquote(input)
	if err != nil {
		return err
	}

	respCards := &sheme.ResponseCards{}
	if err := json.Unmarshal([]byte(jsonInput), respCards); err != nil {
		return err
	}

	fmt.Printf("Your (%s) bank card list:\n", c.User.Login)
	
	for id, card := range respCards.Cards {
		fmt.Printf("%d. PAN: %s; Valid Thru Date: %s; Name: %s; CVV: %s\n", id + 1, card.PAN, card.Name, card.ValidThru, card.CVV)
	}
	
	return nil
}
