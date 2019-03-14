package t3hstr

import "testing"

var gestalttests = []struct {
	name      string
	target    []rune
	maxlen    int
	separator []rune
	expected  string
}{
	{"Odd len", []rune("abcdefg"), 5, []rune("-"), "ab-fg"},
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
