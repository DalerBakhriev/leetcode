package main

import (
	mlocswuc "leetcode/maxlengthofconcatstrwithuniquechars"
	"log"
	"time"
)

func main() {

	start := time.Now()
	uniques := mlocswuc.MaxLength([]string{"cha", "r", "act", "ers"})
	elapsed := time.Since(start)
	log.Println(elapsed)

	log.Println(uniques)

}
