package funcs

import "github.com/shopspring/decimal"

// FuncType - internal type of functions
type FuncType func(args ...decimal.Decimal) (decimal.Decimal, error)

// count of operator priorities
const LevelsOfPriorities = 3
