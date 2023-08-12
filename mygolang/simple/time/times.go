package main

import (
	"fmt"
	"time"
)

func main() {
	times := time.Now()
	fmt.Println(times)
	man_time := times.Format("02-01-2006 15:04:05  Monday")
	fmt.Println(man_time)
}
