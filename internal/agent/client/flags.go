package client

import "flag"

var (
	reg  bool
	auth bool
	file bool

	user     string
	password string
	upload   string
	download string
)

func parseFlags() {
	flag.BoolVar(&reg, "reg", false, "register new user")
	flag.BoolVar(&auth, "auth", false, "authentication/authorization")
	flag.BoolVar(&file, "file", false, "file")

	flag.StringVar(&user, "u", "", "login (use only with flag --reg or --auth)")
	flag.StringVar(&password, "p", "", "password (use only with flag --reg or --auth)")
	flag.StringVar(&upload, "upload", "", "file path to upload to server (use only with flag --file)")
	flag.StringVar(&download, "download", "", "file name to download from server (use only with flag --file)")

	flag.Parse()
}
