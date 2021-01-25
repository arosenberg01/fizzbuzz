package fizzbuzz_test

import (
	"testing"

	"github.com/arosenberg01/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
	}
	actual := fizzbuzz.Do(1, 15)
	if len(expected) != len(actual) {
		t.Errorf("len(fizzbuzz.Do()) = %d, want %d", len(actual), len(expected))
	}
	for i, el := range expected {
		if el != actual[i] {
			t.Errorf("fizzbuzz.Do()[%d] = %s, want %s", i, actual[i], el)
		}
	}
}
