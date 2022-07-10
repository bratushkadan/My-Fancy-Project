package common

import (
	"os"
)

var HmacAuthSecret = []byte(os.Getenv("HMAC_AUTH_SECRET"))

func init() {
	if len(HmacAuthSecret) == 0 {
		panic("Emtpy HmacAuthSecret")
	}
}
