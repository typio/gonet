package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/typio/gonet"
)

// dataToArray processes the text file
func dataToArray(path string) ([][]float64, [][]float64) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	dataInputs := make([][]float64, len(lines))
	dataTargets := make([][]float64, len(lines))
	var s []string
	labels := map[string]float64{"Iris-setosa": 0, "Iris-versicolor": 1 /*, "Iris-virginica": 1*/}
	for i := 0; i < len(lines); i++ {
		s = strings.Split(lines[i], ",")
		dataInputs[i] = make([]float64, len(s)-1)
		dataTargets[i] = make([]float64, 1)
		for j := 0; j < len(s)-1; j++ {
			n, _ := strconv.ParseFloat(s[j], 64)
			dataInputs[i][j] = n
		}
		dataTargets[i][0] = labels[s[len(s)-1]]
	}
	return dataInputs, dataTargets
}

func main() {
	// create 2 2D slices
	dataInputs, dataTargets := dataToArray("./dataShort.txt")

	// Instantiate nn
	nn := gonet.NewNN(4, 4, 1)

	// Train 100000 times
	for i := 0; i < 100000; i++ {
		// select data to train on at random
		rIdx := rand.Intn(100)
		inputs := gonet.FromArray(dataInputs[rIdx])
		targets := gonet.FromArray(dataTargets[rIdx])
		// feed selected input and its target into nn
		nn.Train(inputs, targets)
	}

	// Predict on a Iris-setosa (0)
	fmt.Printf("Should output 0, ouput: %f\n", nn.Predict(gonet.FromArray(dataInputs[0])))
	// Predict on a Iris-versicolor (1)
	fmt.Printf("Should output 1, ouput: %f\n", nn.Predict(gonet.FromArray(dataInputs[50])))
}
