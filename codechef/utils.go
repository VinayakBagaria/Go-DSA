package codechef

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ScanArray() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")

	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers[i] = num
	}
	return numbers
}

func PrintArray(arr []int) string {
	stringNums := make([]string, len(arr))
	for i, num := range arr {
		stringNums[i] = strconv.Itoa(num)
	}

	return strings.Join(stringNums, " ")
}
