package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := readFile("diagnostic-report.txt")
	if err != nil {
		log.Fatal(err)
	}
	// m := make(map[string]int)
	fmt.Println(partOne(lines))
}

func partOne(lines []string) int {
	// fmt.Println(lines)
	var gamma, epsilon string
	for i := 0; i < len(lines[0]); i++ {
		var zero, one int
		for _, binary := range lines {
			if binary[i] == '0' {
				zero++
			} else {
				one++
			}
		}
		if zero > one {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	eps, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}
	gam, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(eps * gam)
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
