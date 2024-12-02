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
	sort.Ints(right)

	var sum int = 0

	for index := range left {
		if left[index]-right[index] < 0 {
			sum += right[index] - left[index]
			fmt.Printf("Subtracting %d from %d\n", left[index], right[index])
			fmt.Printf("The sum is now %d\n", sum)
		} else {
			sum += left[index] - right[index]
			fmt.Printf("Subtracting %d from %d\n", right[index], left[index])
			fmt.Printf("The sum is now %d\n", sum)
		}
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

func splitFile(file *os.File) ([]int, []int) {

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

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
		right = append(right, rint)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}
