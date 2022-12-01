package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var sums []int
	var tot int

	tot = 0
	for (scanner.Scan() ) {
		line := scanner.Text()
		if ( line != "" ) {
			n,_ := strconv.Atoi(line)
			tot += n
		} else {
			sums = append(sums,tot)
			tot = 0
		}
	}
	//fmt.Println(sums)
	max := TopOne(sums)
	fmt.Println("max: ", max)
	three := TopThree(sums)
	fmt.Println("top three: ", three)
}


func TopOne(s []int) int{
	sort.Ints(s)
	//fmt.Println(s)
	return s[len(s)-1]
}

func TopThree(s []int) int{
	sort.Ints(s)
	//fmt.Println(s)
	return s[len(s)-1]+s[len(s)-2]+s[len(s)-3]
}

