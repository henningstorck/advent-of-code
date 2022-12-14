package strutils

import "strconv"

func Atoi(ascii string) int {
	integer, _ := strconv.Atoi(ascii)
	return integer
}
