package fizzbuzz

import "strconv"

// Do gets output of FizzBuzz algorithm from 1-15 inclusive
func Do(n int) []string {
	var out []string
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			out = append(out, strconv.Itoa(i))
		} else {
			out = append(out, s)
		}
	}
	return out
}
