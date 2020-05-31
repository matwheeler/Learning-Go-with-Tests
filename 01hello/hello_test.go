package main

import "testing"

func TestHello(t *testing.T) {

	// We've refactored our assertion into a function.
	// This reduces duplication and improves readability of our tests.
	// In Go you can declare functions inside other functions and assign them to variables.
	// You can then call them, just like normal functions.
	// We need to pass in t *testing.T so that we can tell the test code to fail when we need to.
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		// This will help other developers track down problems easier.
		// If you still don't understand, comment it out, make a test fail and observe the test output.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Here we are introducing another tool in our testing arsenal, subtests. (the t.run())
	// Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
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

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Elodie", "Italian")
		want := "Ciao, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
