package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"time"
)

// Data :nodoc:
type Data struct {
	Int int
}

// ByIntAscSlow sort Data by Int ascendign
type ByIntAscSlow []Data

func (a ByIntAscSlow) Len() int { return len(a) }
func (a ByIntAscSlow) Less(i, j int) bool {
	time.Sleep(time.Nanosecond * 100)
	return a[i].Int < a[j].Int
}
func (a ByIntAscSlow) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return n
}

// current data is int
func populateData() ([]Data, error) {
	file, err := os.Open("data.out")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []Data{}
	for scanner.Scan() {
		data = append(data, Data{Int: stringToInt(scanner.Text())})
	}

	return data, nil
}

func sortNumber(data []Data) {
	sort.Sort(ByIntAscSlow(data))
}