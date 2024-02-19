package day_xx

import (
	"bufio"
	"strings"
)

func solve01(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := ""
	for scanner.Scan() {
		result = scanner.Text()
	}
	return result
}
