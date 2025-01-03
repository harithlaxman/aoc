package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInput() [][]int {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

    var reports [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
        var report []int 
		nums := strings.Split(scanner.Text(), " ")
        for _, v := range nums {
            n, _ := strconv.Atoi(v)
            report = append(report, n)
        }
        reports = append(reports, report)
	}
    return reports
}

func checkArray(arr []int) bool {
	ascendSafe := true
	descendSafe := true
	// Check Ascending
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] || math.Abs(float64(arr[i] - arr[i-1])) > 3 {
			ascendSafe = false
			break
		}
	}
	// Check Descending
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] || math.Abs(float64(arr[i] - arr[i-1])) > 3 {
			descendSafe = false
			break
		}
	}
	return ascendSafe || descendSafe
}

func partOne() {
	count := 0
	reports := getInput()
	for _, arr := range reports {
		if checkArray(arr) {
			count++
		}
	}
	fmt.Println(count)
}

func partTwo() {
	count := 0
	reports := getInput()
	for _, arr := range reports {
		for i := range arr {
			var newArr []int
			for j := range arr {
				if j != i {
					newArr = append(newArr, arr[j])
				}
			}
			if checkArray(newArr) {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

func main() {
    partOne()
	partTwo()
}
