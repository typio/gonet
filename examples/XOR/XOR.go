package main

import (
	"fmt"
	"math/rand"

	"github.com/typio/gonet"
)

func main() {
	dataInputs := [][]float64{{0, 1}, {1, 0}, {0, 0}, {1, 1}}
	dataTargets := [][]float64{{1}, {1}, {0}, {0}}

	nn := gonet.NewNN(2, 2, 1)

	for i := 0; i < 100000; i++ {
		rIdx := rand.Intn(4)
		inputs := gonet.FromArray(dataInputs[rIdx])
		targets := gonet.FromArray(dataTargets[rIdx])
		nn.Train(inputs, targets)
	}

	for i := 0; i < len(dataInputs); i++ {
		fmt.Println(nn.Predict(gonet.FromArray(dataInputs[i])))
	}
}
