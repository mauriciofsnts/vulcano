package documents

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateCPF() string {
	/* #nosec G404 */
	doc := fmt.Sprintf("%v", rand.Float64())[2:11]
	doc += calculateDigit(doc, 10)
	doc += calculateDigit(doc, 11)
	return doc
}

func GenerateCNPJ() string {
	/* #nosec G404 */
	doc := fmt.Sprintf("%v", rand.Float64())[2:10] + "0001"
	doc += calculateDigit(doc, 5)
	doc += calculateDigit(doc, 6)
	return doc
}

func calculateDigit(doc string, position int) string {
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
