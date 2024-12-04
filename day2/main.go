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
	log.Println("Loading input")
	matrix, err := loadMatrix()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Day 2 Part 1")
	numberOfSafeReports := part1(matrix)
	log.Println("Solution = ", numberOfSafeReports)
}

func loadMatrix() ([][]int, error) {
	matrix := make([][]int, 0)
	readFile, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}
	defer readFile.Close()
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		report := make([]int, 0)
		levels := strings.Split(fs.Text(), " ")
		for _, level := range levels {
			conv, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			report = append(report, conv)
		}
		matrix = append(matrix, report)
	}
	return matrix, nil
}

func part1(matrix [][]int) int {
	numOfSafeReports := 0
	for _, report := range matrix {
		numOfSafeReports++
		firstDiff := report[1] - report[0]
		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]
			if abs(diff) > 3 || abs(diff) < 1 || (firstDiff > 0 && diff <= 0) || (firstDiff < 0 && diff >= 0) {
				numOfSafeReports--
				break
			}
		}
	}

	return numOfSafeReports
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
