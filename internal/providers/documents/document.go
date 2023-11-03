package documents

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCPF() (string, string) {
	rand.Seed(time.Now().UnixNano())

	var digits [9]int
	for i := 0; i < 9; i++ {
		digits[i] = rand.Intn(9)
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	digit1 := 11 - (sum % 11)
	if digit1 >= 10 {
		digit1 = 0
	}

	sum = 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (11 - i)
	}
	sum += digit1 * 2
	digit2 := 11 - (sum % 11)
	if digit2 >= 10 {
		digit2 = 0
	}

	withMask := fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d", digits[0], digits[1], digits[2], digits[3], digits[4], digits[5], digits[6], digits[7], digits[8], digit1, digit2)
	withoutMask := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d", digits[0], digits[1], digits[2], digits[3], digits[4], digits[5], digits[6], digits[7], digits[8], digit1, digit2)

	return withMask, withoutMask
}
