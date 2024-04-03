package util

import (
	"bufio"
	"os"
)

func InputLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}
