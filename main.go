package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func calculateReverse(l []int) int64 {
	var matrix [][]int64
	var line []int64
	for i := len(l) - 1; i >= 0; i-- {
		line = append(line, int64(l[i]))
	}
	matrix = append(matrix, line)

	i := 0
	for {
		var sumline int64

		var newLine []int64
		line := matrix[i]
		for w := 0; w < len(line)-1; w++ {
			v := matrix[i][w] - matrix[i][w+1]
			newLine = append(newLine, v)
			sumline += v
		}
		matrix = append(matrix, newLine)
		if sumline == 0 {
			break
		}
		i++
	}

	var newMatrix [][]int64
	var newSum int64
	for w := len(matrix); w > 0; w-- {
		line := matrix[w-1]
		if w == len(matrix) {
			line = append(line, 0)
			newMatrix = append(newMatrix, line)
			i++
			continue
		}
		v := line[len(line)-1]

		linePrec := newMatrix[len(newMatrix)-1]
		diff := linePrec[len(linePrec)-1]

		r := v - diff
		line = append(line, r)
		newMatrix = append(newMatrix, line)
		i++
		if w == 1 {
			newSum += r
		}
	}
	fmt.Println(newMatrix)
	return newSum
}

func calculate(l []int) int {
	var matrix [][]int
	matrix = append(matrix, l)

	i := 0
	for {
		sumline := 0
		var newLine []int
		line := matrix[i]
		for w := 0; w < len(line)-1; w++ {
			v := matrix[i][w+1] - matrix[i][w]
			newLine = append(newLine, v)
			sumline += v
		}
		matrix = append(matrix, newLine)
		if sumline == 0 {
			break
		}
		i++
	}

	var newMatrix [][]int
	newSum := 0
	for w := len(matrix); w > 0; w-- {
		line := matrix[w-1]
		if w == len(matrix) {
			line = append(line, 0)
			newMatrix = append(newMatrix, line)
			i++
			continue
		}
		v := line[len(line)-1]
		linePrec := newMatrix[len(newMatrix)-1]
		diff := linePrec[len(linePrec)-1]
		line = append(line, v-diff)
		newMatrix = append(newMatrix, line)
		i++
		if w == 1 {
			fmt.Printf("%d - %d = %d", v, diff, v-diff)
			newSum += v - diff
		}
	}
	return newSum
}

func main() {
	f, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		var input []int
		for _, l := range strings.Fields(data) {
			i, _ := strconv.Atoi(l)
			input = append(input, i)
		}
		lines = append(lines, input)
	}

	sum := 0
	var sumReverse int64
	exists := make(map[string]int)
	existsReverse := make(map[string]int64)
	for _, l := range lines {
		newSum := 0
		h := sha1.New()
		io.WriteString(h, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(l)), ","), "[]"))

		// Check if we have the same string computed
		v := exists[string(h.Sum(nil))]
		if v != 0 {
			newSum += v
		} else {
			newSum += calculate(l)
		}
		exists[string(h.Sum(nil))] = newSum
		sum += newSum

		var newSumReverse int64
		vr := existsReverse[string(h.Sum(nil))]
		if vr != 0 {
			sumReverse += vr
			continue
		}
		newSumReverse += calculateReverse(l)
		existsReverse[string(h.Sum(nil))] = newSumReverse
		sumReverse += newSumReverse
	}
	fmt.Println(fmt.Sprintf("Sum = %d", sum))
	fmt.Println(fmt.Sprintf("Sum Reverse = %d", sumReverse))
}
