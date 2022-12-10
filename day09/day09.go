package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {

	var n = flag.Int("n", 2, "number of knots in rope")
	flag.Parse()

	var rope []Point
	var visited []Point

	knots := *n
	rope = make([]Point, knots, knots)
	visited = append(visited, Point{0, 0}) // {0,0} is start point

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		word := strings.Split(line, " ")
		dir := word[0]
		dist, _ := strconv.Atoi(word[1])
		for dist > 0 {
			rope[0] = MoveHead(rope[0], dir)
			dist--
			for i := 0; i < len(rope)-1; i++ {
				rope[i+1] = Follow(rope[i], rope[i+1]) // each knot in the rope follows the previous one
			}
			if !Contains(visited, rope[len(rope)-1]) {
				visited = append(visited, rope[len(rope)-1])
			}
		}
	}
	fmt.Println("pos. visited: ", len(visited))
}

func MoveHead(head Point, dir string) Point {

	// move head
	switch dir {
	case "R":
		head.x++
	case "L":
		head.x--
	case "U":
		head.y++
	case "D":
		head.y--
	}
	return head
}

func Follow(leader Point, follower Point) Point {

	var folStep Point

	delta := Sub(leader, follower)

	if (Abs(delta.x) == 2 && (Abs(delta.y) == 0 || Abs(delta.y) == 1)) ||
		(Abs(delta.y) == 2 && (Abs(delta.x) == 0 || Abs(delta.x) == 1)) ||
		(Abs(delta.y) == 2 && Abs(delta.x) == 2) {
		if delta.x != 0 {
			folStep.x = int(math.Copysign(1, float64(delta.x)))
		} else {
			folStep.x = 0
		}
		if delta.y != 0 {
			folStep.y = int(math.Copysign(1, float64(delta.y)))
		} else {
			folStep.y = 0
		}
		follower = Add(follower, folStep)
	}

	return follower
}

func Sub(h Point, t Point) Point {
	return Point{h.x - t.x, h.y - t.y}
}

func Add(h Point, t Point) Point {
	return Point{h.x + t.x, h.y + t.y}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Contains(list []Point, p Point) bool {
	found := false

	for _, v := range list {

		if v == p {
			found = true
			break
		}
	}
	return found
}
