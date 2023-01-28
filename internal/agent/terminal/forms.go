package terminal

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kozyrev-m/keeper/internal/agent/model"
	"github.com/rivo/tview"
)

type Form struct {
	terminal *Terminal

	text *tview.TextView
	formRegister *tview.Form
	formLogin *tview.Form
	formWhoami *tview.Form
}

func NewForm(t *Terminal) *Form {
	f := &Form{
		terminal: t,
		formRegister: tview.NewForm(),
		formLogin: tview.NewForm(),
		formWhoami: tview.NewForm(),
	}

	txt := notauth

	if f.terminal.checkAuth() {
		txt = auth
	}

	f.text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).SetText(txt)

	return f
}

func (f *Form) RegisterUser() {
    user := &model.User{}

    f.formRegister.AddInputField("Login", "", 20, nil, func(login string) {
        user.Login = login
    })

    f.formRegister.AddInputField("Password", "", 20, nil, func(password string) {
        user.Password = password
    })

	f.formRegister.AddButton("Save", func() {
		fmt.Printf("Test: %+v\n", user)
		if err := f.terminal.client.RegisterUser(user); err != nil {
			fmt.Printf("ERROR: %s", err.Error())
		}
		f.terminal.pages.SwitchToPage("Menu")
    })
}

func (f *Form) Login() {
    user := &model.User{}

    f.formLogin.AddInputField("Login", "", 20, nil, func(login string) {
        user.Login = login
    })

    f.formLogin.AddInputField("Password", "", 20, nil, func(password string) {
        user.Password = password
    })

	f.formLogin.AddButton("Send", func() {
		f.text.Clear().SetTextColor(tcell.ColorGreen).
		SetText(auth)

		if err := f.terminal.client.LoginUser(user); err != nil {
			f.text.Clear().SetTextColor(tcell.ColorGreen).
			SetText(notauth).
			SetTextColor(tcell.ColorRed).SetText(err.Error())
		}
		
		f.terminal.pages.SwitchToPage("Menu")
    })
}

func (f *Form) Whoami() {
	u, err := f.terminal.client.Whoami()
	if err != nil {
		f.text.Clear().SetTextColor(tcell.ColorRed).SetText(err.Error())
	}

	txt := fmt.Sprintf("Hello %s (%d)", u.Login, u.ID)
	f.formWhoami.AddTextArea("", txt, 22, 1, len(txt), func(text string) {
		fmt.Printf("Hello %s (%d)", u.Login, u.ID)
	})

	
	// pages.SwitchToPage("Whoami")
}