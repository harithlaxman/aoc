package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput() string {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	var input string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = input + scanner.Text()
	}

	return input
}

func partOne() {
	input := getInput()
	result := 0
	
	for i := 0; i < len(input) - 4; i++ {
		var s string
		for j := i; j < i + 4; j++ {
			s = s + string(input[j])
		}
		if s == "mul(" {
			var num1 int
			num1String := ""
			j := i + 4
			for {
				if string(input[j]) == "," {
					j++
					break
				}
				num1String = num1String + string(input[j])
				j++
			}
			num1, _ = strconv.Atoi(num1String)
			
			var num2 int
			num2String := ""
			for {
				if string(input[j]) == ")" {
					j++
					break
				}
				num2String = num2String + string(input[j])
				j++
			}
			num2, _ = strconv.Atoi(num2String)

			result = result + num1*num2
		}
	}
	fmt.Println(result)
}

func partTwo() {
	do := true

	input := getInput()
	result := 0
	
	for i := 0; i < len(input) - 7; i++ {
		var s string
		
		for j := i; j < i + 7; j++ {
			s = s + string(input[j])
		}
		
		if s == "don't()" {
			do = false
		}
		s = ""
		
		for j := i; j < i + 4; j++ {
			s = s + string(input[j])
		}

		if s == "do()" {
			do = true
		}

		if s == "mul(" && do {
			var num1 int
			num1String := ""
			j := i + 4
			for {
				if string(input[j]) == "," {
					j++
					break
				}
				num1String = num1String + string(input[j])
				j++
			}
			num1, _ = strconv.Atoi(num1String)
			
			var num2 int
			num2String := ""
			for {
				if string(input[j]) == ")" {
					j++
					break
				}
				num2String = num2String + string(input[j])
				j++
			}
			num2, _ = strconv.Atoi(num2String)

			result = result + num1*num2
		}
	}
	fmt.Println(result)
}

func main() {
	partOne()
	partTwo()
}