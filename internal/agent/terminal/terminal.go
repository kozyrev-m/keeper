package terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/kozyrev-m/keeper/internal/agent/httpclient"
)

// Terminal...
type Terminal struct {
	app *tview.Application
	pages *tview.Pages
	form *Form
	client *httpclient.Client
}

// NewTerminal...
func NewTerminal() *Terminal {
	t := &Terminal{
		app: tview.NewApplication(),
		pages: tview.NewPages(),
		client: httpclient.New(),
	}

	t.form = NewForm(t)


	t.initTerminal()

	return t 
}

// initTerminal...
func (t *Terminal) initTerminal() {
	t.pages.AddPage("Menu", t.form.text, true, true)
	t.pages.AddPage("Register User", t.form.formRegister, true, false)
	t.pages.AddPage("Login User", t.form.formLogin, true, false)
	t.pages.AddPage("Whoami", t.form.formWhoami, true, false)


	t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 81 {
			t.app.Stop()
		}
		if event.Rune() == 82 {
			t.form.RegisterUser()
			t.pages.SwitchToPage("Register User")
		}
		if event.Rune() == 83 {
			t.form.Login()
			t.pages.SwitchToPage("Login User")
		}

		if t.checkAuth() {
			if event.Rune() == 87 {
				t.form.Whoami()
				t.pages.SwitchToPage("Whoami")
			}
		}

		return event
	})

}


func (t *Terminal) StartTerminal() error {
	return t.app.SetRoot(t.pages, true).Run()
}

func (t *Terminal) checkAuth() bool {
	
	_, err := t.client.Whoami()
	
	return err == nil
}