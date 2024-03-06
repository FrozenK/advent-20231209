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
	exists := make(map[string]int)
	for _, l := range lines {
		// Create the new matrix
		var matrix [][]int
		matrix = append(matrix, l)

		// Check if we have the same string computed
		h := sha1.New()
		io.WriteString(h, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(l)), ","), "[]"))
		v := exists[string(h.Sum(nil))]
		if v != 0 {
			sum += v
			continue
		}
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
			line = append(line, v+diff)
			newMatrix = append(newMatrix, line)
			i++
			if w == 1 {
				newSum += v + diff
			}
		}
		sum += newSum
		exists[string(h.Sum(nil))] = newSum
		fmt.Println(newMatrix)
	}
	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
