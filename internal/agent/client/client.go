package client

import (
	"errors"

	"github.com/kozyrev-m/keeper/internal/agent/httpclient"
	"github.com/kozyrev-m/keeper/internal/agent/model"
)

func StartClient() error {
	parseFlags()

	client := httpclient.New()
	
	if action == "text" || action == "pair" || action == "card" || action == "file" {
		// check session for --action [text | pair | card | file]
		u, err := client.Whoami()
		if err != nil {
			return err
		}

		client.User = u
	}

	switch action {
	case "register": // register new user in system
		if (len(user) > 0) && (len(password) > 0) {
			u := &model.User{
				Login:    user,
				Password: password,
			}

			return client.RegisterUser(u)
		} else {
			return errors.New("use --action \"register\" with --user and --password")
		}
	case "login": // create session user
		if (len(user) > 0) && (len(password) > 0) {
			u := &model.User{
				Login:    user,
				Password: password,
			}

			return client.LoginUser(u)
		} else {
			return errors.New("use --action \"login\" with --user and --password")
		}
	case "text": // work with text
		if !(len(content) > 0) {
			return client.GetTexts()
		} else {
			txt := &model.Text{
				Metadata: metadata,
				Value:    content,
			}

			return client.AddText(txt)
		}
	case "pair": // work with login-password pairs
		if !(len(user) > 0 || len(password) > 0) {
			return client.GetLoginPasswordPairs()
		} else {
			p := &model.Pair{
				Metadata: metadata,
				Login:    user,
				Password: password,
			}

			return client.AddLoginPasswordPair(p)
		}
	case "card": // work with bank cards
		if !(len(pan) > 0 || len(validThru) > 0 || len(name) > 0 || len(cvv) > 0) {
			return client.GetBankCards()
		} else {
			bc := &model.BankCard{
				Metadata:  metadata,
				PAN:       pan,
				ValidThru: validThru,
				Name:      name,
				CVV:       cvv,
			}

			return client.AddBankCardData(bc)
		}
	case "file": // work with files
		if !(len(upload) > 0 || len(download) > 0) {
			return client.ListFiles()
		}
		if len(upload) > 0 {
			return client.UploadFile(upload, metadata)
		}
		if len(download) > 0 {
			return client.DownloadFile(download)
		}
	default:
		return errors.New("--action is not defined")
	}

	return nil
}
