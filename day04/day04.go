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
	var r1, r2 []int

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		r := strings.Split(line,",")
		r1s := strings.Split(r[0],"-")
		r2s := strings.Split(r[1],"-")
		r1 = nil
		r2 = nil
		r1l,_ := strconv.Atoi(r1s[0])
		r1u,_ := strconv.Atoi(r1s[1])
		r2l,_ := strconv.Atoi(r2s[0])
		r2u,_ := strconv.Atoi(r2s[1])
		r1 = append(r1,r1l)
		r1 = append(r1,r1u)
		r2 = append(r2,r2l)
		r2 = append(r2,r2u)
		if ((r1[0]>=r2[0]) && (r1[1]<=r2[1])) || ((r2[0]>=r1[0]) && (r2[1]<=r1[1])) {
			ncontains++
		}
		if ((r1[1]<r2[0]) || (r1[0]>r2[1]))  {
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


