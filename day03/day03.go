package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	var sum int

	scanner := bufio.NewScanner(os.Stdin)


	for (scanner.Scan() ) {
		var compartment []string
		line := scanner.Text()
		// fmt.Println(line)
		compartment = append(compartment, line[0:len(line)/2])
		compartment = append(compartment, line[len(line)/2:len(line)])
		//fmt.Println(compartment)
		for _, c := range compartment[0] {
			if strings.ContainsRune(compartment[1], c) {
				//fmt.Println(string(c), ":", priority(c))
				sum += priority(c)
				break
			}
		}
	}
	fmt.Println(sum)
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