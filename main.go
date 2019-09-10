package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text := getInput(reader, "Type a comma separated string of items")
	plots := parseInput(text)

	text = getInput(reader, "How many items do you need")
	items := parseInput(text)[0]

	text = getInput(reader, "How many extra items do you need")
	extraItems := parseInput(text)[0]

	randomPlots, cmap := randomSelection(int(items), plots, nil, true)
	extraPlots, _ := randomSelection(int(extraItems), plots, cmap, true)

	fmt.Printf("\n\nWe've selected the following plots:\n%+v\n", randomPlots)
	fmt.Printf("\nWe've selected the following extra plots:\n%v\n", extraPlots)
}

// getInput gets input from the user based on a reader and a question and returns the result as a string
func getInput(r *bufio.Reader, q string) string {
	fmt.Printf("%s:\n", q)
	text, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("an error has occured: %s", err.Error())
		os.Exit(2)
	}
	return text
}

// randomInt returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source based on a seed set by the unix epoch
func randomInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// parseInput cleans and extracts numbers from a string and puts them into an integer array for further processing
func parseInput(input string) []float64 {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\t", "")
	args := strings.Split(input, ",")

	r := make([]float64, 0)

	for _, arg := range args {
		if len(arg) > 0 {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("there was an error converting %s to a number: %s...\n", arg, err.Error())
			}
			r = append(r, num)
		}
	}
	return r
}

// randomSelection generates a random list of items containing n items of the list
func randomSelection(n int, l []float64, c map[int]bool, s bool) ([]float64, map[int]bool) {
	items := make([]float64, n)
	if c == nil {
		c = make(map[int]bool)
	}

	for i := 0; i < n; i++ {
		num := randomInt(0, len(l))
		if c[num] {
			i--
		} else {
			c[num] = true
			items[i] = l[num]
		}
	}

	if s {
		sort.Float64s(items[:])
	}

	return items, c
}
