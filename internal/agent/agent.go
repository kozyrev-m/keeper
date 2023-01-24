package agent

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/kozyrev-m/keeper/internal/agent/httpclient"
	"github.com/kozyrev-m/keeper/internal/agent/model"
)

var app = tview.NewApplication()
var pages = tview.NewPages()
var client = httpclient.New()
/*
var text = tview.NewTextView().
    SetTextColor(tcell.ColorGreen).
    SetText("(q) to quit")
*/

var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(R) register new user \n(S) login \n(Q) quit")

var formRegister = tview.NewForm()
var formLogin = tview.NewForm()
var formWhoami = tview.NewForm()


func StartClient() error {
	configureClient()
	// return app.SetRoot(pages, true).EnableMouse(true).Run()
	return app.SetRoot(pages, true).Run()
}

func configureClient() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 81 {
			app.Stop()
		}
		return event
	})

	pages.AddPage("Menu", text, true, true)
	pages.AddPage("Register User", formRegister, true, false)
	pages.AddPage("Login User", formLogin, true, false)
	pages.AddPage("Whoami", formWhoami, true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 81 {
			app.Stop()
		}
		if event.Rune() == 82 {
			registerUserForm()
			pages.SwitchToPage("Register User")
		}
		if event.Rune() == 83 {
			loginUserForm()
			pages.SwitchToPage("Login User")
		}
		if event.Rune() == 87 {
			whoamiForm()
			pages.SwitchToPage("Whoami")
		}
		return event
	})

}

func registerUserForm() {
    user := &model.User{}

	formRegister.Clear(true)

    formRegister.AddInputField("Login", "", 20, nil, func(login string) {
        user.Login = login
    })

    formRegister.AddInputField("Password", "", 20, nil, func(password string) {
        user.Password = password
    })

	formRegister.AddButton("Save", func() {
		fmt.Printf("Test: %+v\n", user)
		if err := client.RegisterUser(user); err != nil {
			fmt.Printf("ERROR: %s", err.Error())
		}
		pages.SwitchToPage("Menu")
    })
}

func loginUserForm() {
    user := &model.User{}

	formLogin.Clear(true)

    formLogin.AddInputField("Login", "", 20, nil, func(login string) {
        user.Login = login
    })

    formLogin.AddInputField("Password", "", 20, nil, func(password string) {
        user.Password = password
    })

	formLogin.AddButton("Send", func() {
		fmt.Printf("Test: %+v\n", user)
		if err := client.LoginUser(user); err != nil {
			fmt.Printf("ERROR: %s", err.Error())
		}
		text.Clear().SetTextColor(tcell.ColorGreen).
		SetText("(R) register new user \n(S) login \n(W) whoami \n(Q) to quit")
		pages.SwitchToPage("Menu")
    })
}

func whoamiForm() {
	if err := client.Whoami(); err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
}