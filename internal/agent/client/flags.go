package client

import "flag"

var (
	// main flags:
	register bool
	login    bool
	card     bool
	file     bool
	pair     bool
	text     bool

	// additional flags:
	// use with register/login
	user     string
	password string

	// use with file
	upload   string
	download string

	// use with card
	pan       string // PAN (primary account number)
	validThru string
	name      string
	cvv       string // CVV/CVC (Card Verification Value/Code)

	// use with text
	content string

	// metadata
	metadata string
)

func parseFlags() {
	// main flags
	flag.BoolVar(&register, "register", false, "register new user")
	flag.BoolVar(&login, "login", false, "login to your account")
	flag.BoolVar(&text, "text", false, "text")
	flag.BoolVar(&card, "card", false, "bank card")
	flag.BoolVar(&pair, "pair", false, "login-password pair")
	flag.BoolVar(&file, "file", false, "file")

	// additional flags:

	// use with --register or --login or --pair
	flag.StringVar(&user, "u", "", "login (use only with flag --register or --login or --pair)")
	flag.StringVar(&password, "p", "", "password (use only with flag --register or --login or --pair)")

	// use with --register or --login or --pair
	flag.StringVar(&user, "user", "", "user (use only with flag --register or --login or --pair)")
	flag.StringVar(&password, "password", "", "password (use only with --register or --login or --pair)")

	// use with --text
	flag.StringVar(&content, "content", "", "some text (use only with flag --text)")

	// use with --card
	flag.StringVar(&pan, "pan", "", "PAN - primary account number (use only with flag --card)")
	flag.StringVar(&validThru, "till", "", "Valid Thru Date - the expiry date of the card printed on the card surface (use only with flag --card)")
	flag.StringVar(&cvv, "cvv", "", "CVV/CVC - Card Verification Value/Code (use only with flag --card)")
	flag.StringVar(&name, "name", "", "cardholder name - name of the owner, printed on the front of the card (use only with flag --card)")

	// use with --file
	flag.StringVar(&upload, "upload", "", "file path to upload to server (use only with flag --file)")
	flag.StringVar(&download, "download", "", "file name to download from server (use only with flag --file)")

	// --metadata
	flag.StringVar(&metadata, "metadata", "", "metadata -- meta information -- some text")

	flag.Parse()
}
