package main

import (
	"bufio"
	"fmt"
	"strings"
)

type rrr struct {
	a int
}

func (r rrr) Read(p []byte) (n int, err error) {
	return 42, nil
}

// Get first word from stdin
func getFirstWord() string {
	var r rrr
	input := bufio.NewScanner(r)
	input.Scan()
	ans := strings.Fields(input.Text())
	if len(ans) == 0 {
		return ""
	}
	return ans[0]

}
func main() {
	fmt.Printf("Would you like to play a game?\n> ")
	ans := getFirstWord()
	fmt.Printf("Your answer: %s\n", ans)
}
