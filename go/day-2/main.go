package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := readFile("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(partOne(file))
	fmt.Println(partTwo(file))
}

func partOne(input []string) int {
	horizontal := 0
	depth := 0
	for i := 0; i < len(input); i++ {
		direction := input[i]
		directionChange, err := strconv.Atoi(strings.Fields(direction)[1])
		if err != nil {
			log.Fatal()
		}
		switch {
		case strings.Contains(direction, "forward"):
			horizontal += directionChange

		case strings.Contains(direction, "down"):
			depth += directionChange

		case strings.Contains(direction, "up"):
			depth -= directionChange
		}

	}
	return horizontal * depth
}

func partTwo(input []string) int {
	horizontal := 0
	depth := 0
	aim := 0
	for i := 0; i < len(input); i++ {
		direction := input[i]
		directionChange, err := strconv.Atoi(strings.Fields(direction)[1])
		if err != nil {
			log.Fatal()
		}
		switch {
		case strings.Contains(direction, "forward"):
			horizontal += directionChange
			if aim > 0 {
				depth += aim * directionChange
			}

		case strings.Contains(direction, "down"):
			aim += directionChange

		case strings.Contains(direction, "up"):
			aim -= directionChange
		}

	}
	return horizontal * depth
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			return nil, err
		}
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
