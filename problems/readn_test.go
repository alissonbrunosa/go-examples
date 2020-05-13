package problems

import (
	"testing"
)

type tts struct {
	value    int
	expected string
}

func testReadN(t *testing.T, st *StringBuffer, tts []tts) {
	for _, tt := range tts {
		if got := st.ReadN(tt.value); got != tt.expected {
			t.Errorf("Expected '%s', got '%s'", tt.expected, got)
		}
	}
}

func TestReadN(t *testing.T) {
	r := NewStringBuffer("Hello World")

	testReadN(t, r, []tts{
		{7, "Hello W"},
		{7, "orld"},
		{7, ""},
		{7, ""},
	})

	r = NewStringBuffer("Hello World")
	testReadN(t, r, []tts{
		{13, "Hello World"},
		{13, ""},
	})
}
