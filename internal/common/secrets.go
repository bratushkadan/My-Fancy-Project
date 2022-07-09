package common

import (
	"os"
)

var HmacAuthSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
