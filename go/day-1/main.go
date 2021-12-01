package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readFile("sea-floor-measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	partOne(input)
	fmt.Printf("Part one: %d\n", partOne(input))
	fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input []int) int {
	largerThanPreviousIterCount := 0
	for idx, seaFloorDepth := range input {
		if idx != 0 && input[idx-1] < seaFloorDepth {
			largerThanPreviousIterCount += 1
		}
	}
	return largerThanPreviousIterCount
}

func partTwo(input []int) int {
	largerThanPreviousThreeIterCount := 0
	currSeaFloorDepth := 100000000000

	for i := 0; i < len(input)-2; i++ {
		curSumOf := input[i] + input[i+1] + input[i+2]
		if curSumOf > currSeaFloorDepth {
			largerThanPreviousThreeIterCount += 1
		}
		currSeaFloorDepth = curSumOf

	}
	return largerThanPreviousThreeIterCount
}

func readFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed_line, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, parsed_line)
	}

	return lines, scanner.Err()
}
