package client

import "flag"

var (
	reg      bool
	auth     bool
	user     string
	password string
)

func parseFlags() {
	flag.BoolVar(&reg, "reg", false, "register new user")
	flag.BoolVar(&auth, "auth", false, "authentication/authorization")
	flag.StringVar(&user, "u", "", "login (use only with flag --reg or --auth)")
	flag.StringVar(&password, "p", "", "password (use only with flag --reg or --auth)")

	flag.Parse()
}
