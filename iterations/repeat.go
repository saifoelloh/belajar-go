package iterations

import (
	"fmt"
	"strings"
)

var repeatedCount = 5

func Repeat(character string) string {
	var repeated strings.Builder

	for range repeatedCount {
		repeated.WriteString(character)
	}

	return repeated.String()
}

func main() {
	fmt.Println("hello world")
}
