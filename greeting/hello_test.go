package greeting

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		// given (preconditions):
		want := "Hello, Chris!"

		// when (action):
		got := Hello("Chris", "English")

		// then (expected result):
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		// given (preconditions):
		want := "Hello, world!"

		// when (action):
		got := Hello("", "English")

		// then (expected result):
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		// given (preconditions):
		want := "Hola, Elodie!"

		// when (action):
		got := Hello("Elodie", "Spanish")

		// then (expected result):
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		// given (preconditions):
		want := "Bonjour, Yosuke!"

		// when (action):
		got := Hello("Yosuke", "French")

		// then (expected result):
		assertCorrectMessage(t, got, want)
	})
}
