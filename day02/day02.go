package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var score, total, thand, mhand int

	score = 0
	total = 0

	for (scanner.Scan() ) {
		line := scanner.Text()
		hand := strings.Split(line, " ")
//		fmt.Println("orig: ",hand)

		// adjust second col to acieve win/loss/draw
		switch hand[1] {
			case "X":  //should win
				hand[1] =  SelectLoser(hand[0]) 
			case "Y":  //should lose
				hand[1] =  SelectTie(hand[0]) 
			case "Z":  //should draw
				hand[1] =  SelectWinner(hand[0]) 
		}
//		fmt.Println("trans: ", hand)
		

		switch hand[0] {
			case "A":
				thand = 1
			case "B":
				thand = 2
			case "C":
				thand = 3
		}
		score = 0
		switch hand[1] {
			case "X":
				mhand = 1
				score += 1
			case "Y":
				mhand = 2
				score += 2
			case "Z":
				mhand = 3
				score += 3
		}
		total += score + ScoreBout(thand,mhand)
	}
	fmt.Println("score: ", total)
}

func ScoreBout(them int, me int) int {

	var score int

	val := me - them
	//fmt.Println("val: ",  val)
	if val == 0 {
		score = 3
	} else if ( val == 1 || val == -2 ) {
		score = 6
	} else {
		score = 0
	}
	//fmt.Println("bout score: ", score)
	return score

}

func SelectWinner(them string) string {

	var hand string

	switch them {
		case "A":
			hand = "Y"
		case "B":
			hand = "Z"
		case "C":
			hand = "X"
	}
//	fmt.Println("return: ", hand)
	return hand
}

func SelectLoser(them string) string {

	var hand string
	switch them {
		case "A":
			hand = "Z"
		case "B":
			hand = "X"
		case "C":
			hand = "Y"
	}
//	fmt.Println("return: ", hand)
	return hand

}

func SelectTie(them string) string {

	var hand string
	switch them {
		case "A":
			hand = "X"
		case "B":
			hand = "Y"
		case "C":
			hand = "Z"
	}
//	fmt.Println("return: ", hand)
	return hand

}

