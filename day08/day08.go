package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


func main() {

	var nrows, ncols int
	var irow int
	var maprow []int
	var treemap [][]int

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		nrows++
		ncols = len(line)
//		fmt.Println(line)
		maprow = nil
		for _, c := range line {
//			fmt.Println(string(c))
			height, _ := strconv.Atoi(string(c))			
			maprow = append(maprow, height)
		}
		treemap = append(treemap,maprow)
		irow++
	}
	fmt.Println(nrows,ncols)
	//fmt.Println(treemap)
	nvis := 0
	maxscore := 0
	score := 0
	for r,_ := range treemap {
		for c,_ := range treemap[r] {
			if IsVisible(r,c,treemap) {
				nvis++
			}
			score = ScenicScore(r,c,treemap)
			if score>maxscore {
				 maxscore= score
			}
		}
	}
	fmt.Println(nvis)
	fmt.Println(maxscore)
}

func ScenicScore(r int, c int, tmap [][]int) int {

	var uscore, dscore, lscore, rscore int

	//fmt.Println("r",r,"c",c,"len(tmap)",len(tmap),"len(tmap[r])",len(tmap[r]))
	if r == 0 || c == 0 || r == len(tmap) || c == len(tmap[r]) {
		// add 0
	}
	//fmt.Println("calculating scenic score for  tree at ",r," ",c," with height ",tmap[r][c])
	// right
	//fmt.Println("right")
	for col:=c+1; col<len(tmap[r]); col++ {
		//fmt.Println("checking against tree at ",r," ",col," with height ",tmap[r][col])
		if tmap[r][col] >= tmap[r][c] {
			rscore++
			break
		} else {
			rscore++
		}
	}
	//fmt.Println("rscore ",rscore)
	// left
	//fmt.Println("left")
	for col:=c-1; col>=0; col-- {
		//fmt.Println("checking against tree at ",r," ",col," with height ",tmap[r][col])
		if tmap[r][col] >= tmap[r][c] {
			lscore++
			break
		} else {
			lscore++
		}
	}
	//fmt.Println("lscore ", lscore)
	// up
	//fmt.Println("up")
	for row:=r-1; row>=0; row-- {
		//fmt.Println("checking against tree at ",row," ",c," with height ",tmap[row][c])
		if tmap[row][c] >= tmap[r][c] {
			uscore++
			break
		} else {
			uscore++
		}
	}
	//fmt.Println("uscore ", uscore)
	// down
	//fmt.Println("down")
	for row:=r+1; row<len(tmap); row++ {
		//fmt.Println("checking against tree at ",row," ",c," with height ",tmap[row][c])
		if tmap[row][c] >= tmap[r][c] {
			dscore++
			break
		} else {
			dscore++
		}
	}
	//fmt.Println("dscore ", dscore)
	return uscore*dscore*lscore*rscore
}

func IsVisible(r int, c int, tmap [][]int) bool {

	var isVis bool = true
	var isVisR, isVisL, isVisU, isVisD = true, true, true, true
	//fmt.Println("r",r,"c",c,"len(tmap)",len(tmap),"len(tmap[r])",len(tmap[r]))
	if r == 0 || c == 0 || r == len(tmap) || c == len(tmap[r]) {
		return isVis
	}
	//fmt.Println("testing if tree at ",r," ",c," with height ",tmap[r][c]," is visible")
	// right
	//fmt.Println("right")
	for col:=c+1; col<len(tmap[r]); col++ {
		//fmt.Println("checking against tree at ",r," ",col," with height ",tmap[r][col])
		if tmap[r][col] >= tmap[r][c] {
			//fmt.Println("it blocks it!")
			isVisR = false
		}
	}
	// left
	//fmt.Println("left")
	for col:=c-1; col>=0; col-- {
		//fmt.Println("checking against tree at ",r," ",col," with height ",tmap[r][col])
		if tmap[r][col] >= tmap[r][c] {
			//fmt.Println("it blocks it!")
			isVisL = false
		}
	}
	// up
	//fmt.Println("up")
	for row:=r-1; row>=0; row-- {
		//fmt.Println("checking against tree at ",row," ",c," with height ",tmap[row][c])
		if tmap[row][c] >= tmap[r][c] {
			//fmt.Println("it blocks it!")
			isVisU = false
		}
	}
	// down
	//fmt.Println("down")
	for row:=r+1; row<len(tmap); row++ {
		//fmt.Println("checking against tree at ",row," ",c," with height ",tmap[row][c])
		if tmap[row][c] >= tmap[r][c] {
			//fmt.Println("it blocks it!")
			isVisD = false
		}
	}
	isBlocked := !isVisR && !isVisL && !isVisU && !isVisD
	return !isBlocked
}