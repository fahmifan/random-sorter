package main

import (
	"bufio"
	"os"
	"strconv"
)

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return n
}

// current data is int
func populateData(fpath string) ([]int, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []int
	for scanner.Scan() {
		data = append(data, stringToInt(scanner.Text()))
	}

	return data, nil
}

// buble sorts
func sortNumber(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				// swap
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
