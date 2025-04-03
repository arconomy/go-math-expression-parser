package basic

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/arconomy/go-math-expression-parser/funcs"
	"github.com/shopspring/decimal"
)

var (
	// the array of operations sorted by operators
	// operators[0] - highest operators (unary, functions)
	// operators[1] - medium operators (*, /, %, ^)
	// operators[2] - lowest operators (+, -)
	DefaultOperators = [funcs.LevelsOfPriorities]map[string]funcs.FuncType{
		{
			"+":    UnarySum,
			"-":    UnarySub,
			"sqrt": Sqrt,
			"abs":  Abs,
		},
		{
			"*": Mult,
			"/": Div,
			"%": DivReminder,
			"^": Pow,
		},
		{
			"+": Sum,
			"-": Sub,
		},
	}
)

func UnarySum(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 1 {
		return decimal.Zero, errors.New("incorrect count of args for unary sum operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0], nil
}
func UnarySub(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 1 {
		return decimal.Zero, errors.New("incorrect count of args for unary subtract operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0].Neg(), nil
}

func Sqrt(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 1 {
		return decimal.Zero, errors.New("incorrect count of args for 'sqrt' function. Need: 1, but get: " + strconv.Itoa(len(args)))
	}
	if args[0].LessThan(decimal.Zero) {
		return decimal.Zero, errors.New("'sqrt' function argument is negative: " + fmt.Sprintf("%v", args[0]))
	}
	return decimal.NewFromFloat(math.Sqrt(args[0].InexactFloat64())), nil
}

func Abs(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 1 {
		return decimal.Zero, errors.New("incorrect count of args for 'abs' function. Need: 1, but get: " + strconv.Itoa(len(args)))
	}
	return args[0].Abs(), nil
}

func Mult(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("incorrect count of args for multiplication operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0].Mul(args[1]), nil
}

func Div(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("incorrect count of args for division operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	if args[1].IsZero() {
		return decimal.Zero, errors.New("incorrect divisor for division operator")
	}

	return args[0].Div(args[1]), nil
}

func Pow(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("incorrect count of args for power operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0].Pow(args[1]), nil
}

func DivReminder(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("incorrect count of args for % operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	if args[1].IsZero() {
		return decimal.Zero, errors.New("incorrect divisor for % operator")
	}
	return args[0].Mod(args[1]), nil
}

func Sum(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("incorrect count of args for sum operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0].Add(args[1]), nil
}

func Sub(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("incorrect count of args for subtract operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0].Sub(args[1]), nil
}
