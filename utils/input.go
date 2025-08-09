package utils

import (
	"bufio"
	"os"
)

var Scanner = bufio.NewScanner(os.Stdin)

func ReadLine() string {
	Scanner.Scan()
	return Scanner.Text()
}
