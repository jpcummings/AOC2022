package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
	"flag"
)

type Monkey struct {
	items []int
	op string
	// e1 "string"  e1 is always "old"
	e2 string
	test int
	truetarget int
	falsetarget int
	ninspected int
}

var nrounds *int
var relaxfac *int

func main() {
	var monkeys []Monkey

	nrounds= flag.Int("n",20,"number of rounds")
	relaxfac = flag.Int("r",3,"relaxation factor")
	flag.Parse()

	scanner:= bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if words[0] == "Monkey" {
			var m Monkey
			scanner.Scan()
			line = strings.Trim(scanner.Text()," ")
			words = strings.Split(line, " ")
			for i:=2; i<len(words); i++ {
				item,_ := strconv.Atoi(strings.Split(words[i],",")[0])
				m.items = append(m.items, item)
			}
			scanner.Scan()
			line = strings.Trim(scanner.Text()," ")
			words = strings.Split(line, " ")
			m.op = words[4]
			m.e2 = words[5]
			scanner.Scan()
			line = strings.Trim(scanner.Text()," ")
			words = strings.Split(line, " ")
			m.test,_ = strconv.Atoi(words[3])
			scanner.Scan()
			line = strings.Trim(scanner.Text()," ")
			words = strings.Split(line, " ")
			m.truetarget,_ = strconv.Atoi(words[5])
			scanner.Scan()
			line = strings.Trim(scanner.Text()," ")
			words = strings.Split(line, " ")
			m.falsetarget,_ = strconv.Atoi(words[5])

			monkeys = append(monkeys, m)
		}
	}
	for iround := 1; iround < (*nrounds+1); iround++ {
		monkeys = DoRound(monkeys)
	}
	fmt.Println(MonkeyBusiness(monkeys))
}


func DoRound(monkeys []Monkey) []Monkey {
	var worry int
	for imonkey, m := range monkeys {
		for _, item := range m.items {
			monkeys[imonkey].ninspected++
			worry = WorryUpdate(item, m.op, m.e2)/(*relaxfac)
			if worry%m.test == 0 {
			// throw to truetarget
				monkeys[m.truetarget].items = append(monkeys[m.truetarget].items, worry%(17*7*13*2*19*3*5*11))
			} else {
			// throw to falsetarg
				monkeys[m.falsetarget].items = append(monkeys[m.falsetarget].items, worry%(17*7*13*2*19*3*5*11))
			}
			// remove from m
			monkeys[imonkey].items = monkeys[imonkey].items[1:]
		}
	}
	return monkeys
}

func WorryUpdate(item int, op string, e2 string) int {

	var arg, ret int

	if e2 == "old" {
		arg = item
	} else {
		arg,_ = strconv.Atoi(e2)
	}

	switch op {
	case "*":
		ret = item*arg
	case "+":
		ret = item+arg
	}
	return ret
}

func MonkeyBusiness(monks []Monkey) int {
	sort.Slice(monks, func(i, j int) bool {
		return monks[i].ninspected < monks[j].ninspected
	})
	return monks[len(monks)-1].ninspected*monks[len(monks)-2].ninspected
}

