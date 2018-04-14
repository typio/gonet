# GoNet 

### A Simple (*incomplete*) Machine Learning Framework and Matrix Library
For the purpose of explaining a neural network from scratch

*a lot of this is based on [this](https://github.com/CodingTrain/Toy-Neural-Network-JS/tree/master/lib)*

# Getting Started
### Install GoNet:
```
$ go get github.com/typio/gonet
```

# Documentation
### Basics:
* Import the package:  
    `import "github.com/typio/gonet"`

### Matrix Library:
* Creates a 2 by 3 matrix of zeros:  
    `matrix := gonet.Create(2, 3)`

* Returns the matrix:  
    `matrix.Read()`