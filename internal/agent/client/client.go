package client

import (
	"errors"

	"github.com/kozyrev-m/keeper/internal/agent/httpclient"
	"github.com/kozyrev-m/keeper/internal/agent/model"
)

func StartClient() error {
	parseFlags()

	client := httpclient.New()

	// register new user in system
	if register && !(len(user) > 0 && len(password) > 0) {
		return errors.New("use flag --register with --user and --password")
	}
	if register && (len(user) > 0) && (len(password) > 0) {
		u := &model.User{
			Login:    user,
			Password: password,
		}

		return client.RegisterUser(u)
	}

	// create session user
	if login && !(len(user) > 0 && len(password) > 0) {
		return errors.New("use flag --login with --user and --password")
	}
	if login && (len(user) > 0) && (len(password) > 0) {
		u := &model.User{
			Login:    user,
			Password: password,
		}

		return client.LoginUser(u)
	}

	// check session
	u, err := client.Whoami()
	if err != nil {
		return err
	}
	client.User = u

	// text
	if text && (len(content) > 0) {
		txt := &model.Text{
			Metadata: metadata,
			Value:    content,
		}

		return client.AddText(txt)
	}
	if text && !(len(content) > 0) {
		return client.GetTexts()
	}

	// login-password pair
	if pair && (len(user) > 0 || len(password) > 0) {
		p := &model.Pair{
			Metadata: metadata,
			Login:    user,
			Password: password,
		}

		return client.AddLoginPasswordPair(p)
	}
	if pair && !(len(user) > 0 || len(password) > 0) {
		return client.GetLoginPasswordPairs()
	}

	// bank card data
	if card && (len(pan) > 0 || len(validThru) > 0 || len(name) > 0 || len(cvv) > 0) {
		bc := &model.BankCard{
			Metadata:  metadata,
			PAN:       pan,
			ValidThru: validThru,
			Name:      name,
			CVV:       cvv,
		}

		return client.AddBankCardData(bc)
	}
	if card && !(len(pan) > 0 || len(validThru) > 0 || len(name) > 0 || len(cvv) > 0) {
		return client.GetBankCards()
	}

	// list/upload/download file
	if file && !(len(upload) > 0 || len(download) > 0) {
		return client.ListFiles()
	}
	if file && len(upload) > 0 {
		return client.UploadFile(upload, metadata)
	}
	if file && len(download) > 0 {
		return client.DownloadFile(download)
	}

	return nil
}
