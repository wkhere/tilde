package tilde

import (
	"testing"
)

func TestExpand(t *testing.T) {
	home, err := homeFromEnv()
	if err != nil {
		panic(err)
	}

	var tab = []struct{ input, want string }{
		{"", ""},
		{"a", "a"},
		{"a/", "a/"},
		{"/", "/"},
		{"/foo", "/foo"},
		{"/foo/", "/foo/"},
		{"~", home},
		{"~/", home + "/"},
		{"~/dir", home + "/dir"},
		{"~/dir/x", home + "/dir/x"},
		{"~/~", home + "/~"},
		{"~/a~", home + "/a~"},
		{"~/a/~", home + "/a/~"},
		{"~a", "~a"},
		{"/~", "/~"},
		{"/~/", "/~/"},
		{"/~/q", "/~/q"},
	}

	for i, tc := range tab {
		have, err := Expand(tc.input)
		if err != nil {
			t.Errorf("tc#%d unexpected error %s", i, err)
		}
		if have != tc.want {
			t.Errorf("tc#%d have %q, want %q", i, have, tc.want)
		}
	}
}
