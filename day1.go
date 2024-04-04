package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ch1(scanner bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		l, r := -1, -1

		for _, b := range scanner.Bytes() {
			if b < 97 {
				if l == -1 {
					l, _ = strconv.Atoi(string(b))
					r, _ = strconv.Atoi(string(b))
				}
				r, _ = strconv.Atoi(string(b))
			}
		}

		sum = sum + 10*l + r
	}
	return sum
}

func ch2(scanner bufio.Scanner) int {
	sum := 0
	arr := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		l, r := -1, -1
		b := scanner.Bytes()
		for i, by := range b {
			if by < 97 {
				if l == -1 {
					l = i
					r = i
				}
				r = i
			}
		}
		ls := string(b[0:l])
		rs := string(b[r:])
		li := -1
    ll := 0
		ri := -1
    rr := 0
		for j, s := range arr {
			i := strings.Index(ls, s)
			if i != -1 {
				if li == -1 {
					li = i
					ll = j + 1
				}
				if i < li {
					li = i
					ll = j + 1
				}
			}
			i = strings.Index(rs, s)
			if i != -1 {
				if ri == -1 {
					ri = i
					rr = j + 1
				}
				if i > ri {
					ri = i
					rr = j + 1
				}
			}
		}
    if ll != 0 {
      l = ll 
    } else {
      l, _ = strconv.Atoi(string(b[l]))
    }
    if rr != 0 {
      r = rr
    } else {
      r, _ = strconv.Atoi(string(b[r]))
    }
    sum = sum + 10*l + r
	}

	return sum
}

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println(ch2(*scanner))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
