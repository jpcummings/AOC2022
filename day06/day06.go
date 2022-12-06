package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {


	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		matchlen := 14
		for i  := 0; i < len(line)-(matchlen-1); i++ {
			if Unique(line[i:i+matchlen]) {
				fmt.Println("testing ", line[i:i+matchlen])
				fmt.Println("unique at ", i+matchlen)
				break
			}
		}
	}
}

func Unique(s string) bool {
	var tested string = ""
	for _, c := range s {
		if strings.Contains(tested, string(c)) {
			return false
		} else {
			tested += string(c)
		}
	}
	return true
}