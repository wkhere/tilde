package tilde

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Expand(name string) (_ string, err error) {
	name = filepath.Clean(name)
	sep := string(filepath.Separator)

	pp := strings.SplitN(name, sep, 2)
	if pp[0] == "~" {
		v, err := homeFromEnv()
		if err != nil {
			return name, fmt.Errorf("expand %s: %w", name, err)
		}
		pp[0] = filepath.Clean(v)
	}
	return strings.Join(pp, sep), nil
}
