package util

import "os"

func DefaultEnv(envVar string, defaultVal string) string {
	if val := os.Getenv(envVar); val != "" {
		return val
	}
	return defaultVal
}
