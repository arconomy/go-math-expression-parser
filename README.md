# expp - tiny math expression parser

[![Go Report Card](https://goreportcard.com/badge/github.com/arconomy/go-math-expression-parser)](https://goreportcard.com/report/github.com/arconomy/go-math-expression-parser)
[![Coverage Status](https://coveralls.io/repos/github/Overseven/go-math-expression-parser/badge.svg?branch=main)](https://coveralls.io/github/Overseven/go-math-expression-parser?branch=main)

## Changes
This is a fork of https://github.com/Overseven/go-math-expression-parser. The main difference is that `shopspring.Decimal` is used for precision.

## Contents
- [expp - tiny math expression parser](#expp---tiny-math-expression-parser)
  - [Changes](#changes)
  - [Contents](#contents)
  - [Supported operations](#supported-operations)
  - [Example](#example)
  - [User-defined functions](#user-defined-functions)
  - [TODO](#todo)

## Supported operations
This parser supports some elements of math expressions:
- unary operators `+, -`
- binary operators `+, -, *, /, ^, %`
- any variables without spaces and operator symbols
- parenthesis `10*(x%(4+y))`
- functions `sqrt(x), abs(x)`
- user defined functions with a comma-separated list of arguments
 
## Example
This part contains the example of parsing and evaluating expression:
```go
s := "(price - purchasePrice) * numOfGoods * 0.87"
```

Create `expp.Parser` object:
```go
parser := expp.NewParser()
```


To parse expression call `parser.Parse()` function. `expp.Exp` string conversation returns string with [prefix style operation notation](http://www.cs.man.ac.uk/~pjj/cs212/fix.html) 
```go
exp, _ := parser.Parse(s)
fmt.Println("Parsed execution tree: ", exp)
// Parsed execution tree: ( * ( * ( - price purchasePrice ) numOfGoods ) 0.87 )
```

To get sorted list of all variables used in the expression call ``expp.GetVarList()`` function:
```go
vars := expp.GetVarList(exp)
fmt.Println("Variables: ", vars)
// Variables: [numOfGoods price purchasePrice]
```
All variables must be defined to calculate an expression result:
```go
values := make(map[string]decimal.Decimal)
values["numOfGoods"] = 20
values["price"] = 15.4
values["purchasePrice"] = 10.3
``` 
Getting the result of evaluation:
```go
result, _ := parser.Evaluate(values)
fmt.Println("Result: ", result)
// Result: 88.74
```
The additional example is contained in the `console_calc.go` [file](https://github.com/arconomy/go-math-expression-parser/blob/main/console_calc.go)

## User-defined functions
You can add to the parser your own function and set the expression string presentation name.
To do this, you need to create `expp.Parser` object with using `expp.NewParser` function
```go
package main

import (
	"fmt"

	"github.com/arconomy/go-math-expression-parser/expp"
)

// Foo - example of user-defined function
func Foo(a ...decimal.Decimal) (decimal.Decimal, error) {
	fmt.Println("Foo was called!")
	var sum decimal.Decimal
	for _, val := range a {
		sum = sum.Add(val)
	}
	return sum, nil
}

func main() {
    s := "10 * bar(60, 6, 0.6)"
    
    // create parser object
    parser := expp.NewParser()
    
    // add function to parsing
    parser.AddFunction(Foo, "bar")
    
    // parsing
    exp, err := parser.Parse(s)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    
    fmt.Println("\nParsed execution tree:", exp)
    // output: 'Parsed execution tree: ( * 10 ( bar ( 60,6,0.6 ) ) )'
    
    // execution of the expression
    result, err := parser.Evaluate(map[string]decimal.Decimal{})
    if err != nil {
        fmt.Println("Error: ", err)
    }
    
    fmt.Println("Result: ", result)
    // output: 'Result: 666' 
}
```
## TODO
- [x] binary operators 
- [x] unary operators
- [x] simple predefined functions (like `sqrt(x)` and `abs(x)`)
- [x] comma-separated list of arguments
- [x] [user-defined functions](#user-defined-functions)
- [x] tests
- [x] create struct `expp.Parser`, which contains parser context with included user-defined functions  
