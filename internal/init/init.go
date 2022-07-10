package init

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	handleEnv()
}

func handleEnv() {
	if buildEnv := os.Getenv("BUILD_ENV"); buildEnv == "" || buildEnv == "testing" {
		godotenv.Load()
	}
}
