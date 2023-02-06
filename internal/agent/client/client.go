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
	if reg && !(len(user) > 0 && len(password) > 0) {
		return errors.New("use flag --reg with -u and -p")
	}
	if reg && (len(user) > 0) && (len(password) > 0) {
		u := &model.User{
			Login:    user,
			Password: password,
		}

		return client.RegisterUser(u)
	}

	// create session user
	if auth && !(len(user) > 0 && len(password) > 0) {
		return errors.New("use flag --auth with -u and -p")
	}
	if auth && (len(user) > 0) && (len(password) > 0) {
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
			Value: content,
		}

		return client.AddText(txt)
	}
	if text && !(len(content) > 0) {
		return client.GetTexts()
	}

	// login-password pair
	if pair && (len(login) > 0 || len(password) > 0) {
		p := &model.Pair{
			Login: login,
			Password: password,
		}

		return client.AddLoginPasswordPair(p)
	}
	if pair && !(len(login) > 0 || len(password) > 0) {
		return client.GetLoginPasswordPairs()
	}

	// bank card data
	if card && (len(pan) > 0 || len(validThru) > 0 || len(name) > 0 || len(cvv) > 0) {
		bc := &model.BankCard{
			PAN: pan,
			ValidThru: validThru,
			Name: name,
			CVV: cvv,
		}

		return client.AddBankCardData(bc)
	}
	if card && !(len(pan) > 0 || len(validThru) > 0 || len(name) > 0 || len(cvv) > 0) {
		return client.GetBankCards()
	}

	// list/upload/download file
	if file && !(list || len(upload) > 0 || len(download) > 0) {
		return errors.New("use flag --file with --list or --upload or --download")
	}
	if file && len(upload) > 0 {
		return client.UploadFile(upload, metadata)
	}
	if file && len(download) > 0 {
		return client.DownloadFile(download)
	}
	if file && list {
		return client.ListFiles()
	}

	return nil
}
