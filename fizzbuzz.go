package fizzbuzz

import "strconv"

// Run start the fizz buzz test wich count from 1 to 100 include in i
// if i is multiple of n1, so print v1
// is i is multiple of n2, so print v2
// else jsut print number
func Run(n1, n2 int, v1, v2 string) string {
	var result string

	if n1 == 0 || n2 == 0 {
		return "error"
	}
	for i := 1; i <= 100; i++ {
		flag := false
		if i%n1 == 0 {
			result += v1
			flag = true
		}
		if i%n2 == 0 {
			result += v2
			flag = true
		}
		if !flag {
			result += strconv.Itoa(i)
		}
		if i < 100 {
			result += ";"
		}
	}
	return result
}
