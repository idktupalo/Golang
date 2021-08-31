package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := ReadStr()
	words := strings.Fields(phrase)
	abbr := ""
	for _, word := range words {
		firstLetter := []rune(word)[0]
		if unicode.IsLetter(firstLetter) {
			abbr += string(unicode.ToUpper(firstLetter))
		}
	}
	fmt.Println(abbr)
}

func ReadStr() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
