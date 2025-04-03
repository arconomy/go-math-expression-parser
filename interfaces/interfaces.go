package interfaces

import (
	"github.com/arconomy/go-math-expression-parser/funcs"
	"github.com/shopspring/decimal"
)

type ExpParser interface {
	AddFunction(f funcs.FuncType, s string)
	GetFunctions() [funcs.LevelsOfPriorities]map[string]funcs.FuncType
	String() string
	Parse(str string) (Expression, error)
	Evaluate(vars map[string]decimal.Decimal) (decimal.Decimal, error)
}

// Exp - the base interface for Term and Node structures
type Expression interface {
	String() string
	Evaluate(vars map[string]decimal.Decimal, p ExpParser) (decimal.Decimal, error)
	GetVarList(vars map[string]interface{})
}

// Function - the struct which contains a function and an argument
type Function interface {
	Expression
	SetOperation(string)
	GetOperation() string
	SetArgs([]Expression)
	GetArgs() []Expression
}
