package utils

import (
	"testing"
)

func TestReverseName(t *testing.T) {
	t.Run("ReverseString", func(t *testing.T) {
		got := ReverseString("Hello, world")
		want := "dlrow ,olleH"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
