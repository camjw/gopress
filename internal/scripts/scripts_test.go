package scripts

import "testing"

var regextests = []struct {
	in      string
	diffs   []byte
	regexes []string
	out     bool
}{
	{"match one regex", []byte("account_page.js"), []string{"page"}, true},
	{"match multiple regexes", []byte("123.js"), []string{"notmatching", "\\d+"}, true},
	{"doesn't match regex", []byte("testing_page.js"), []string{"login"}, false},
}

func TestCheckRegexesAgainstDiffs(t *testing.T) {
	for _, tt := range regextests {
		t.Run(tt.in, func(t *testing.T) {
			match := CheckRegexesAgainstDiffs(tt.diffs, tt.regexes)
			if match != tt.out {
				t.Errorf("got %v, want %v, expressions: %s, diffs: %s", match, tt.out, tt.regexes, string(tt.diffs))
			}
		})
	}
}
