package internal

import (
	"errors"
	"strconv"

	"github.com/arconomy/go-math-expression-parser/interfaces"
	"github.com/shopspring/decimal"
)

// Term - the struct which contains a single value
type Term struct {
	Val string
}

func (t *Term) GetVarList(vars map[string]interface{}) {
	if t.Val == "" {
		return
	}
	if _, err := strconv.ParseFloat(t.Val, 64); err == nil {
		return
	}
	vars[t.Val] = struct{}{}

}

// Evaluate - return a value which contains in Term
func (t *Term) Evaluate(vars map[string]decimal.Decimal, p interfaces.ExpParser) (decimal.Decimal, error) {
	if t.Val == "" {
		return decimal.Zero, nil
	}
	if val, err := decimal.NewFromString(t.Val); err == nil {
		return val, nil
	}
	val, ok := vars[t.Val]
	if !ok {
		return decimal.Zero, errors.New("value '" + t.Val + " not found in map")
	}
	return val, nil
}

// toString conversation
func (t *Term) String() string {
	return t.Val
}
