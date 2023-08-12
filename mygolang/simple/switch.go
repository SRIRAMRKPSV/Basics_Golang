// UPTO LINE "16" ITS BELONGS TO RANDOM NUMBER

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("random number Program")
	rand.Seed(time.Now().UnixNano())
	dicenum := rand.Intn(6) + 1
	fmt.Println(dicenum)

	//switch case

	switch dicenum {
	case 1:
		fmt.Println("open the quota..if already opened then move one step")
	case 2:
		fmt.Println("move two steps")
	case 3:
		fmt.Println("move three steps")
	case 4:
		fmt.Println("move four steps")
	case 5:
		fmt.Println("move five steps and roll it again")
	case 6:
		fmt.Println("move six steps and roll it again")
	default:
		fmt.Println("there is any error")
	}

}
