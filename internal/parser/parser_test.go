package parser

import (
	"testing"
)

type mockReader struct {
	s string
}

func (m mockReader) Read(b []byte) (n int, err error) {
	n = copy(b, m.s)
	return n, nil
}

var configtests = []struct {
	label string
	in    string
	out   string
}{
	{"trailing slash", "{\"directory\":\"first/\",\"extension\":\".test\",\"basebranch\":\"origin/master\",\"tests\":[{\"testfile\":\"test\",\"regexes\":[\".*\"]}]}", "first/test.test"},
	{"no trailing slash", "{\"directory\":\"second/\",\"extension\":\".feature\",\"basebranch\":\"origin/master\",\"tests\":[{\"testfile\":\"test\",\"regexes\":[\".*\"]}]}", "second/test.feature"},
}

func TestGetFilePath(t *testing.T) {
	for _, tt := range configtests {
		t.Run(tt.label, func(t *testing.T) {
			reader := mockReader{tt.in}
			config, err := GetConfig(reader)
			if err != nil {
				t.Error("Error occurred during test: ", err)
			}

			got := config.GetFilePath(config.Tests[0])
			want := tt.out
			if got != want {
				t.Errorf("Incorrect filepath. Got: %s. Want: %s.", got, want)
			}
		})
	}
}
