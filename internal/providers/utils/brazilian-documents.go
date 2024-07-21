package utils

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// this is just for development purposes
// the documents doesn't really exist or are valid in any way
func GenerateCPF() string {
	/* #nosec G404 */
	doc := fmt.Sprintf("%v", rand.Float64())[2:11]
	doc += calculateDigits(doc, 10)
	doc += calculateDigits(doc, 11)
	return doc
}

func GenerateCNPJ() string {
	/* #nosec G404 */
	doc := fmt.Sprintf("%v", rand.Float64())[2:10] + "0001"
	doc += calculateDigits(doc, 5)
	doc += calculateDigits(doc, 6)
	return doc
}

func calculateDigits(doc string, position int) string {
	var sum int
	for _, r := range doc {
		sum += int(r-'0') * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
