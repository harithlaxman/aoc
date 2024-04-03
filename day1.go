package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ch1(scanner bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		var n []byte

		for _, b := range scanner.Bytes() {
			if b < 97 {
				n = append(n, b)
			}
		}

		s := string(n[0]) + string(n[len(n)-1])
		no, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		if no < 10 {
			no = no*10 + no
		}

		sum += no
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
  
  fmt.Println(ch1(*scanner))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
