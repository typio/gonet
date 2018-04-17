package gonet

import "math"

// NeuralNetwork is the structure of the network
type NeuralNetwork struct {
	inputNodes, hiddenNodes, outputNodes int
	weightsIH, weightsHO                 *Matrix
	biasH, biasO                         *Matrix
	learningRate                         float64
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

	learningRate := .1

	return &NeuralNetwork{inputNodes, hiddenNodes, outputNodes,
		weightsIH, weightsHO,
		biasH, biasO,
		learningRate}
}

// Sigmoid is the logistic function
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// DSigmoid is the derivative of the logistic function and is run after the Sigmoid func
func DSigmoid(y float64) float64 {
	// return Sigmoid(x) * (1 - Sigmoid(x))
	return y * (1 - y)
}

// Predict gets output for input
func (nn *NeuralNetwork) Predict(inputs *Matrix) float64 {
	// Generates hidden outputs
	hidden := nn.weightsIH.MatrixP(inputs)
	hidden.AddM(nn.biasH)
	hidden.MapM(Sigmoid)

	outputs := nn.weightsHO.MatrixP(hidden)
	outputs.AddM(nn.biasO)
	outputs.MapM(Sigmoid)

	return outputs.data[0][0]
}

// Train uses backprop to modify weights
func (nn *NeuralNetwork) Train(inputs *Matrix, targets *Matrix) {
	hidden := nn.weightsIH.MatrixP(inputs)
	hidden.AddM(nn.biasH)
	hidden.MapM(Sigmoid)

	outputs := nn.weightsHO.MatrixP(hidden)
	outputs.AddM(nn.biasO)
	outputs.MapM(Sigmoid)

	// Calculate error
	// ERROR = TARGETS - OUTPUTS
	outputErrors := targets.SubtractM(outputs)

	// Calculate gradient
	gradients := outputs.MapNM(DSigmoid)
	gradients.MultiplyM(outputErrors)
	gradients.Multiply(nn.learningRate)

	// Calculate deltas
	hiddenT := hidden.Transpose()
	weightsHOD := gradients.MatrixP(hiddenT)

	// Adjust weights by deltas
	nn.weightsHO.AddM(weightsHOD)
	// Adjusts the bias by gradients
	nn.biasO.AddM(gradients)

	// Calculate hidden layer errors
	weightsHOT := nn.weightsHO.Transpose()
	hiddenErrors := weightsHOT.MatrixP(outputErrors)

	// Calculate hidden gradient
	hiddenGradient := hidden.MapNM(DSigmoid)
	hiddenGradient.MultiplyM(hiddenErrors)
	hiddenGradient.Multiply(nn.learningRate)

	// Calculate input -> deltas
	inputsT := inputs.Transpose()
	weightsIHD := hiddenGradient.MatrixP(inputsT)

	nn.weightsIH.AddM(weightsIHD)
	// Adjusts the bias by gradients
	nn.biasH.AddM(hiddenGradient)
}
