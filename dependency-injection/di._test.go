package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Frank")

	got := buffer.String()
	want := "Hello, Frank"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
