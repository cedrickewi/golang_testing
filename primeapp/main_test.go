package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 6, false, "6 is not a prime number because is divisible by 2!"},
		{"not prime", -61, false, "-61 is not prime because it is negative"},
		{"not prime", 0, false, "0 is not a prime, by def"},
		{"not prime", 1, false, "1 is not a prime, by def"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_promt(t *testing.T) {
	// safe a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writter
	_ = w.Close()

	// reset os.Stdout
	os.Stdout = oldOut

	// read the out of our promt func from our read pipe

	out, _ := io.ReadAll(r)

	// perform our test

	if string(out) != "->" {
		t.Errorf("incorrect prompt: expected -> but got %s", out)
	}

}

func Test_intro(t *testing.T) {
	oldstdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	// reset os.Stdout
	os.Stdout = oldstdout

	// read the out of our promt func from our read pipe

	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Is it Prime?") {
		t.Errorf("Intro text not correct: Expected: 'Is it Prime?', Got: %s", out)
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not a prime, by def"},
		{name: "one", input: "1", expected: "1 is not a prime, by def"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "negative", input: "-41", expected: "-41 is not prime because it is negative"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s ", e.name, e.expected, res)
		}
	}

}
