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

	log.Printf("session for '%s' was created successfully!\n", u.Login)

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

// UploadFile sents file to server.
func (c *Client) UploadFile(filepath, metadata string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", path.Base(filepath))
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.WriteField("metadata", metadata); err != nil {
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

	if resp.StatusCode != 200 {
		return errors.New(c.error(resp.Body))
	}

	log.Println("File uploaded!")

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

	if resp.StatusCode != 200 {
		return errors.New(c.error(resp.Body))
	}

	// Create the file
	filepath := fmt.Sprintf("/tmp/%s", filename)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	log.Println("File downloaded!")

	return nil
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

	fmt.Printf("Your (%s) files:\n", c.User.Login)
	for i, file := range respFiles.Files {
		fmt.Println()
		fmt.Printf("%d. :/%s [meta information: \"%s\"]\n", i+1, file.Name, file.Metadata)
	}

	return nil
}

// AddBankCardData adds bank card data.
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

// GetBankCards gets bank card list.
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
		fmt.Println()
		fmt.Printf(
			"%d. PAN: '%s'; Valid Thru Date: '%s'; Name: '%s'; CVV: '%s'\n[meta information: \"%s\"]\n",
			id+1, card.PAN, card.Name, card.ValidThru, card.CVV, card.Metadata,
		)
	}

	return nil
}

// AddLoginPasswordPair adds login-password pair.
func (c *Client) AddLoginPasswordPair(p *model.Pair) error {
	b, err := c.encoder(p)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/pair", http.MethodPost, b)
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

	log.Println("Add login-password pair!")
	return nil
}

// GetLoginPasswordPairs gets login-password pairs.
func (c *Client) GetLoginPasswordPairs() error {
	b, err := c.encoder(nil)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/pair", http.MethodGet, b)
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

	respPairs := &sheme.ResponsePairs{}
	if err := json.Unmarshal([]byte(jsonInput), respPairs); err != nil {
		return err
	}

	fmt.Printf("Your (%s) login-password pairs:\n", c.User.Login)

	for id, pair := range respPairs.Pairs {
		fmt.Println()
		fmt.Printf("%d. Login: '%s'; Password: '%s' [meta information: \"%s\"]\n", id+1, pair.Login, pair.Password, pair.Metadata)
	}

	return nil
}

// AddText adds some text.
func (c *Client) AddText(txt *model.Text) error {
	b, err := c.encoder(txt)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/text", http.MethodPost, b)
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

	log.Println("Add text!")
	return nil
}

// GetTexts gets texts.
func (c *Client) GetTexts() error {
	b, err := c.encoder(nil)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("/private/text", http.MethodGet, b)
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

	respTexts := &sheme.ResponseTexts{}
	if err := json.Unmarshal([]byte(jsonInput), respTexts); err != nil {
		return err
	}

	fmt.Printf("Your (%s) texts:\n", c.User.Login)

	for id, text := range respTexts.Texts {
		fmt.Println()
		fmt.Printf("%d. \"%s\" [meta information: \"%s\"]\n", id+1, text.Value, text.Metadata)
	}

	return nil
}
