package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	log.Println("Loading input")
	list1, list2, err := loadLists()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Day 1 Part 1")
	result, err := part1(list1, list2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Result 1 = ", result)

	log.Println("Day 1 Part 2")
	result, err = part2(list1, list2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Result 2 = ", result)
}

func loadLists() ([]int, []int, error) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	readFile, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("could not read file: %w", err)
	}
	defer readFile.Close()
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		parts := strings.Split(fs.Text(), "   ")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("wrong input format. Like should contain 2 elements but was: %v", len(parts))
		}
		err = convertAndAddToList(parts[0], &list1)
		if err != nil {
			return nil, nil, err
		}
		err = convertAndAddToList(parts[1], &list2)
		if err != nil {
			return nil, nil, err
		}
	}

	return list1, list2, nil
}

func convertAndAddToList(element string, list *[]int) error {
	conv, err := strconv.Atoi(element)
	if err != nil {
		return err
	}
	*list = append(*list, conv)
	return nil
}

func part1(list1 []int, list2 []int) (int, error) {
	size := len(list1)
	if size != len(list2) {
		return -1, errors.New("list sizes are not equal")
	}
	slices.Sort(list1)
	slices.Sort(list2)

	sumOfDistances := 0
	for i := 0; i < size; i++ {
		sumOfDistances += abs(list1[i] - list2[i])
	}
	return sumOfDistances, nil
}

func part2(list1 []int, list2 []int) (int, error) {
	numOfOccurancesByNumber := make(map[int]int)
	for _, el := range list2 {
		numOfOccurancesByNumber[el]++
	}
	similiarityScore := 0
	for _, el := range list1 {
		numOfOccurances, found := numOfOccurancesByNumber[el]
		if found {
			similiarityScore += el * numOfOccurances
		}
	}
	return similiarityScore, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
