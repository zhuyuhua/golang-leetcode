package main

import (
	"fmt"
	"golang-leetcode/calc"
)


const (
	INDEX_COMPOSED = iota
	INDEX_ENTRY
	INDEX_XIAOQIANG
	INDEX_BANNER
	INDEX_NEWS
	INDEX_RESEARCH
)

func main() {

	fmt.Println(INDEX_BANNER)

	fmt.Println(INDEX_NEWS)

	calc.CalcTest()
	
}
