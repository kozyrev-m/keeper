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
	_, err := client.Whoami()
	if err != nil {
		return err
	}

	// upload file
	if file && !(len(filepath) > 0) {
		return errors.New("use flag --file with --upload")
	}
	if file && len(filepath) > 0 {
		return client.UploadFile(filepath)
	}

	return nil
}
