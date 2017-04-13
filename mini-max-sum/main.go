package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(l string) []int {

	s := strings.Split(l, " ")
	n := make([]int, len(s))

	for i, v := range s {
		c, _ := strconv.Atoi(v)
		n[i] = c
	}

	return n
}

type IntSlice []int

func remove(i int, p []int) []int {
	// test := append(p[:i], p[i+1:]...) // perfectly fine if i is the last element
	return append(p[:i], p[i+1:]...)
}

func sum(l []int) int {
	fmt.Println("data:", l)
	s := 0
	for _, v := range l {
		s += v
	}
	return s
}

func min(l []int) int {
	m := l[0]
	for _, v := range l {
		if m > v {
			m = v
		}
	}
	return m
}

func max(l []int) int {
	m := l[0]
	for _, v := range l {
		if m < v {
			m = v
		}
	}
	return m
}

func main() {
	r := bufio.NewReader(os.Stdin)
	l, _ := r.ReadString('\n')
	in := parseInput(l)
	fmt.Println(in)

	sums := make([]int, len(in))

	for i, _ := range in {
		fmt.Println(i)
		sums[i] = sum(remove(i, parseInput(l)))
		// sums[i] = sum(remove(i, in))
	}

	fmt.Println(min(sums), max(sums))
}
