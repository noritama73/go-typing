package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var STRS []string
var LIMIT_TIME_BY_SECOND time.Duration = 90

func main() {
	t := make(chan time.Time)
	fin := make(chan string)
	STRS = []string{"あしたもいいてんき", "きょうのごはんはかれー"}
	go checkStrings(fin)
	countTime(t, fin)
}

func checkStrings(fin chan<- string) {
	for itr := 0; itr < len(STRS); itr++ {
		fmt.Println("No.", itr, " : ", STRS[itr])
		for recieveStdin() != STRS[itr] {
			fmt.Println("WRONG...")
		}
		fmt.Println("CORRECT!")
	}
	fin <- "GAME CLEAR!!"
}

func recieveStdin() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	if s.Err() != nil {
		panic("")
	}
	return s.Text()
}

func countTime(t <-chan time.Time, f <-chan string) {
	select {
	case res := <-f:
		fmt.Println(res)
	case <-time.After(LIMIT_TIME_BY_SECOND * time.Second):
		fmt.Println("Time Over...")
	}
}
