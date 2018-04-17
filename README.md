# GoNet

### A Simple Machine Learning Framework and Matrix Library
For the purpose of explaining a neural network from scratch

# Getting Started
### Install GoNet:
```
$ go get github.com/typio/gonet
```

# Documentation
### Basics:
* Import the package:  
    `import "github.com/typio/gonet"`

### Neural Network Library:

* Instantiates a neural network with 4 input nodes, 4 hidden nodes, and 1 output node:  
    `nn := gonet.NewNN(4, 4, 1)`

* Trains the network once on matrix inputsM and matrix targetsM:  
    `nn.Train(inputsM, targetsM)`

* Returns a guess on 1 by n matrix; representing a data point, as a float64:  
    `nn.Predict(gonet.FromArray(dataInputs[0]))`

### Matrix Library:

#### Basic Methods
* Creates a 2 by 3 matrix of zeros:  
    `m := gonet.Create(2, 3)`

* Creates a 1 by n matrix from a 1D slice:  
    `m := gonet.fromArray([]float)`

* Returns the matrix:  
    `m.Read()`

* Prints the matrix and its dimensions:  
    `m.PrintM()`

* Returns int array of matrix's dimensions (*[rows, cols]*):  
    `m.GetSize()`

* Fills matrix with random floats in range [-1, 1):  
    `m.Randomize()`

#### Scalar Methods
* Adds n (*float64*) to every element in matrix:  
    `m.Add(n)`

* Multiplies n (*float64*) to every element in matrix:  
    `m.Multiply(n)`

#### Elementwise Methods
* Returns new matrix of which every element is m[i][j] + n[i][j] (*must both be same dimensions*):  
    `s := m.AddM(n)`

* Returns new matrix of which every element is m[i][j] - n[i][j] (*must both be same dimensions*):  
    `d := m.SubtractM(n)`

* Returns new matrix of which every element is m[i][j] * n[i][j]  (*must both be same dimensions*):  
    `p := m.MultiplyM(n)`

* Passes every element in matrix through func fn():  
    `p := m.MapM(fn)`

* Returns new matrix in which every element in matrix is passed through func fn()  :  
    `p := m.MapNM(fn)`

#### Linear Algebra Methods
* Returns new matrix which is the product of matrix m and matrix n (*dimensions of m and n must have m cols = n rows*)  
    `p := m.MatrixP(n)`

* Returns new transposed matrix (*swaps rows and cols*)  
    `mT := m.Tranpose()`