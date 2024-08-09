package consoleReader

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func ReadString(message string) (string) {
	fmt.Print(message)
	var reader = bufio.NewReader(os.Stdin)
	var input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Print("oops! Somthing get wrong!")
		return ""
	}
	return input
}


func ReadNumber(message string) (float64) {
	var input = ReadString(message)
	input = strings.TrimSpace(input)
	var num, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Print("oops! Somthing get wrong!")
		return -1
	}
	return num
}


func main() {
	fmt.Printf("Пипи")
}
