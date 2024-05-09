// B4sbT5
package FizzBuzz

import (
	"fmt"
)

func answer(n int) (number int, msg string) {
	msg = ""
	if n%3 == 0 {
		msg = "Fizz"
	}
	if n%5 == 0 {
		msg += "Buzz"
	}
	number = n
	return
}

func main() {
	i := 0
	for i < 99 {
		number, msg := answer(i)
		i += 1
		if number%3 == 0 || number%5 == 0 {
			fmt.Println(msg)
		} else {
			fmt.Println(number)
		}
	}
}
