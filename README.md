# GoNet (*incomplete*)

### A Simple Machine Learning Framework and Matrix Library
For the purpose of explaining a neural network from scratch

*a lot of this is based on [this](https://github.com/CodingTrain/Toy-Neural-Network-JS/tree/master/lib)*

# Getting Started
### Install GoNet:
```
$ go get github.com/typio/gonet
```

# Documentation (*very incomplete*)
### Basics:
* Import the package:  
    `import "github.com/typio/gonet"`

### Matrix Library:

#### Basic Methods
* Creates a 2 by 3 matrix of zeros:  
    `m := gonet.Create(2, 3)`

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
* Adds corresponding element in matrix n to matrix m (*must both be same dimensions*):  
    `m.AddM(n)`

* Adds corresponding element in matrix n to matrix m (*must both be same dimensions*):  
    `m.MultiplyM(n)`

#### Linear Algebra Methods
* Returns new matrix which is the product of matrix m and matrix n (*dimensions of m and n must have [rows, cols] = [cols, rows]*)  
    `p := m.MatrixP(n)`

* Returns new transposed matrix (*swaps rows and cols*)  
    `mT := m.Tranpose()`