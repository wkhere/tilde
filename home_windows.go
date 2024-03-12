package tilde

import "os"

func homeFromEnv() (string, error) {
	if v := os.Getenv("HOME"); v != "" {
		return v, nil
	}
	return os.UserHomeDir()
}
