package helloworld

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("sayng hello to people", func(t *testing.T) {
		t.Run("in Spanish", func(t *testing.T) {
			got := Hello("Frank", "Spanish")
			want := "Hola, Frank"

			assertCorrectMessage(t, got, want)
		})

		t.Run("in French", func(t *testing.T) {
			got := Hello("Frank", "French")
			want := "Bonjour, Frank"

			assertCorrectMessage(t, got, want)
		})

	})
	t.Run("say 'Hello, World' when an empty string is suppied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})
}
