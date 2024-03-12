package tilde

import "os"

func homeFromEnv() (string, error) {
	return os.UserHomeDir()
}
