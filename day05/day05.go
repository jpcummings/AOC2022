package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Mode int

const (
	header Mode = iota
	moves
)

var stack [9][]string

func main() {

	var mode Mode = header

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		if (line == "") {
			mode = moves
			scanner.Scan()
			line = scanner.Text()
		}
		if (mode == header) {
			//fmt.Println(line)
			for i, c := range line {
				if (string(c) == "[") {
					stack[i/4] = append(stack[i/4],string(line[i+1]))
				}
			}
		}
		if (mode == moves) {
			words := strings.Split(line, " ")
			nmoves, _ := strconv.Atoi(words[1])
			origstack, _ := strconv.Atoi(words[3])
			deststack, _ := strconv.Atoi(words[5])
			Move9001(nmoves, origstack, deststack)
		}
	}
	PrintTops()
}

func Move(movesleft int, orig int, dest int) {
	for ( movesleft > 0 ) {
		//fmt.Println("move",movesleft,"from", orig, "to", dest)
		stack[dest-1] = append([]string{stack[orig-1][0]},stack[dest-1]...)
		stack[orig-1] = stack[orig-1][1:]
		movesleft--
	}
	// PrintStacks()
}

func Move9001(nmoves int, orig int, dest int) {
	pick := make([]string, len(stack[orig-1][0:nmoves]))
	copy(pick, stack[orig-1][0:nmoves])
	stack[dest-1] = append(pick, stack[dest-1]...)
	stack[orig-1] = stack[orig-1][nmoves:]
	//PrintStacks()
}

func PrintStacks() {
	for _, s := range stack {
		fmt.Println(s)
	}
}

func PrintTops() {
	var ret string = ""
	for _, s := range stack {
		if s!=nil {
			//fmt.Println(s[0])
			ret += s[0]
		}
	}
	fmt.Println(ret)
}
