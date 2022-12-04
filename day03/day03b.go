package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	var lineno, sum int
	var elf []string

	scanner := bufio.NewScanner(os.Stdin)


	for (scanner.Scan() ) {
		line := scanner.Text()
		lineno++
		// fmt.Println(line)
		elf = append(elf, line)

		if lineno%3 == 0 {
			//fmt.Println(elf)
			//fmt.Println("")
			for _, c := range elf[0] {
				if strings.ContainsRune(elf[1], c) && strings.ContainsRune(elf[2], c) {
					//fmt.Println(string(c), ":", priority(c))
					sum += priority(c)
					break
				}
			}
			elf = nil
		}
	}
	fmt.Println("sum: ", sum)
}

func priority(r rune) int {
	var p rune 
	if r > 64 && r < 91 {  //Uppercase
		p = r - 64 + 26
	} else if r > 96 && r < 123 {  // Lowercase
		p = r - 96
	} else {
		fmt.Println("unknown character: ", r)
		os.Exit(1)
	}
	return int(p)
}