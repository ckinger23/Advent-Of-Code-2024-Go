package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := openFile()
	defer file.Close()

	left, right := splitFile(file)

	sort.Ints(left)

	var sum int = 0

	for _, num := range left {
		sum += num * right[num]
		fmt.Printf("The sum is now %d\n", sum)
	}

	fmt.Println(sum)
}

func openFile() *os.File {

	file, err := os.OpenFile("./d1-input.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func splitFile(file *os.File) ([]int, map[int]int) {

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Fields(line)

		lint, err := strconv.Atoi(split_line[0])
		if err != nil {
			log.Fatal(err)
		}

		rint, err := strconv.Atoi(split_line[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, lint)
		right[rint] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}
