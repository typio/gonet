package gonet

import "math"

// NeuralNetwork is the structure of the network
type NeuralNetwork struct {
	inputNodes, hiddenNodes, outputNodes int
	weightsIH, weightsHO                 *Matrix
	biasH, biasO                         *Matrix
}

// NewNN initializes network
func NewNN(inputNodes, hiddenNodes, outputNodes int) *NeuralNetwork {
	weightsIH := Create(hiddenNodes, inputNodes)
	weightsHO := Create(outputNodes, hiddenNodes)
	weightsIH.Randomize()
	weightsHO.Randomize()

	biasH := Create(hiddenNodes, 1)
	biasO := Create(outputNodes, 1)
	biasH.Randomize()
	biasO.Randomize()

	return &NeuralNetwork{inputNodes, hiddenNodes, outputNodes,
		weightsIH, weightsHO,
		biasH, biasO}
}

// Sigmoid is the logistic function
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// FeedForward does the processess of the network
func (nn *NeuralNetwork) FeedForward(inputs *Matrix) *Matrix {
	// Generates hidden outputs
	hidden := nn.weightsIH.MatrixP(inputs)
	hidden.AddM(nn.biasH)
	hidden.MapM(Sigmoid)

	output := nn.weightsHO.MatrixP(hidden)
	output.MapM(Sigmoid)

	return output
}

// // Train uses backprop to modify weights
// func (nn *NeuralNetwork) Train(inputs *Matrix, targets *Matrix) {
// 	outputs := nn.FeedForward(inputs)

// 	// Calculate error
// 	// ERROR = TARGETS - OUTPUTS
// 	error := targets.SubtractM(outputs)
// }
