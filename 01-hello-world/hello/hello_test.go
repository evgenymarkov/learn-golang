package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Evgeny", "English")
		want := "Hello, Evgeny"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Russian", func(t *testing.T) {
		got := Hello("Евгений", "Russian")
		want := "Привет, Евгений"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Evgeny", "French")
		want := "Bonjour, Evgeny"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Привет, мир' when empty string is supplied in Russian", func(t *testing.T) {
		got := Hello("", "Russian")
		want := "Привет, мир"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Bonjour, monde' when empty string is supplied in Russian", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, monde"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
