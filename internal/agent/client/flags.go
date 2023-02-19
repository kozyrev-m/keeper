package client

import "flag"

var (
	// main flags:
	action string

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
	flag.StringVar(&action, "action", "", "--action [register | login | text | card]")

	// additional flags:
	//  use with --action "register", "login" or "pair"
	flag.StringVar(&user, "u", "", "login (use only with --action \"register\", \"login\" or \"pair\")")
	flag.StringVar(&password, "p", "", "password (use only with --action \"register\", \"login\" or \"pair\")")

	//  use with --action "register", "login" or "pair"
	flag.StringVar(&user, "user", "", "user (use only with --action \"register\", \"login\" or \"pair\")")
	flag.StringVar(&password, "password", "", "password (use only with --action \"register\", \"login\" or \"pair\")")

	//  use with --action "text"
	flag.StringVar(&content, "content", "", "some text (use only with --action \"text\")")

	//  use with --action "card"
	flag.StringVar(&pan, "pan", "", "PAN - primary account number (use only with --action \"card\")")
	flag.StringVar(&validThru, "till", "", "Valid Thru Date - the expiry date of the card printed on the card surface (use only with --action \"card\")")
	flag.StringVar(&cvv, "cvv", "", "CVV/CVC - Card Verification Value/Code (use only with --action \"card\")")
	flag.StringVar(&name, "name", "", "cardholder name - name of the owner, printed on the front of the card (use only with --action \"card\")")

	//  use with --action "file"
	flag.StringVar(&upload, "upload", "", "file (set path of file) to upload to server (use only with --action \"file\")")
	flag.StringVar(&download, "download", "", "file (set name of file) to download from server (use only with with --action \"file\")")

	// --metadata - optionally
	flag.StringVar(&metadata, "metadata", "", "metadata - meta information; some additional text")

	flag.Parse()
}
