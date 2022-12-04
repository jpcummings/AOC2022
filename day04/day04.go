package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	var ncontains,noverlap int

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		r := strings.Split(line,",")
		r1s := strings.Split(r[0],"-")
		r2s := strings.Split(r[1],"-")
		r1l,_ := strconv.Atoi(r1s[0])
		r1u,_ := strconv.Atoi(r1s[1])
		r2l,_ := strconv.Atoi(r2s[0])
		r2u,_ := strconv.Atoi(r2s[1])
		if ((r1l>=r2l) && (r1u<=r2u)) || ((r2l>=r1l) && (r2u<=r1u)) {
			ncontains++
		}
		if ((r1u<r2l) || (r1l>r2u))  {
			// disjoint
		} else {
			// overlap
			//fmt.Println(line)
			noverlap++
		}
	}
	fmt.Println("ncontains: ",ncontains)
	fmt.Println("noverlap: ",noverlap)
}


