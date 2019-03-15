package t3hstr

import "testing"

var gestalttests = []struct {
	name      string
	target    []rune
	maxlen    int
	separator []rune
	expected  string
}{
	{"Odd len odd separator", []rune("abcdefghijk"), 7, []rune("-"), "abc-ijk"},
	{"Odd len even separator", []rune("abcdefghijk"), 7, []rune("--"), "ab--jk"},
	{"Even len odd separator", []rune("abcdefghijk"), 8, []rune("-"), "abc-ijk"},
	{"Even len even separator", []rune("abcdefghijk"), 8, []rune("--"), "abc--ijk"},
	{"Text shorter than maxlen", []rune("abcdefghijk"), 20, []rune("--"), "abcdefghijk"},
	{"Text same len as maxlen", []rune("abcdefghijk"), 11, []rune("--"), "abcdefghijk"},
	{"Text one longer than maxlen", []rune("abcdefghijk"), 10, []rune("--"), "abcd--hijk"},
}

func TestGestalt(t *testing.T) {
	for _, tt := range gestalttests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Gestalt(tt.target, tt.maxlen, tt.separator)

			if actual != tt.expected {
				t.Errorf("got %s, want %s", actual, tt.expected)
			}
		})
	}
}
