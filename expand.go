package tilde

import (
	"fmt"
	"os"
	"strings"
)

func Expand(path string) (string, error) {
	home, sep, err := readEnv()
	if err != nil {
		return path, fmt.Errorf("expand %s: %w", path, err)
	}

	pp := strings.SplitN(path, sep, 2)
	if pp[0] == "~" {
		pp[0] = home
	}
	return strings.Join(pp, sep), nil
}

func readEnv() (home, sep string, _ error) {
	if v := os.Getenv("HOME"); v != "" {
		return v, "/", nil
	}
	sep = string(os.PathSeparator)
	v, err := os.UserHomeDir()
	return v, sep, err
}
