package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getArrays() ([]int, []int) {
	file, err := os.Open("day1_1.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	var leftArr []int
	var rightArr []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		left_num, _ := strconv.Atoi(nums[0])
		right_num, _ := strconv.Atoi(nums[1])
		leftArr = append(leftArr, left_num)
		rightArr = append(rightArr, right_num)
	}
	return leftArr, rightArr
}

func partOne() {
	leftArr, rightArr := getArrays()
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	totalDistance := 0
	for i := 0; i < len(leftArr); i++ {
		totalDistance += int(math.Abs(float64(leftArr[i]) - float64(rightArr[i])))
	}

	fmt.Println(totalDistance)
}

func partTwo() {
	leftArr, rightArr := getArrays()

	totalScore := 0

	for i := 0; i < len(leftArr); i++ {
		count := 0
		for j := 0; j < len(rightArr); j++ {
			if leftArr[i] == rightArr[j] {
				count++
			}
		}
		totalScore += leftArr[i] * count
	}

	fmt.Println(totalScore)
}

func main() {
	partOne()
	partTwo()
}
