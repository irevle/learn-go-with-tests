// Writing a test is just like writing a function, with a few rules
//
//	It needs to be in a file with a name like xxx_test.go
//	The test function must start with the word Test
//	The test function takes one argument only t *testing.T
//	To use the *testing.T type, you need to import "testing", like we did with fmt in the other file
package main

import "testing"

func TestHello(t *testing.T) {
	// t.Run for having subtests in a test
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Elv", "")
		want := "Hello, Elv"

		// %q for passing variable inside formatted Error
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// t = test
// b = benchmark
// t.Helper() tells the test suite that this method is a helper
// got string, want string can be shortened to got, want string (because both has same type)
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
