package internal_test

import (
	"strconv"
	"testing"

	"github.com/arconomy/go-math-expression-parser/internal"
	expp "github.com/arconomy/go-math-expression-parser/parser"
	"github.com/shopspring/decimal"
)

const float64EqualityThreshold = 1e-9

func fuzzyEqual(a, b decimal.Decimal) bool {
	return a.Sub(b).Abs().LessThanOrEqual(decimal.NewFromFloat(float64EqualityThreshold))
}

func TestEvalWithVars(t *testing.T) {
	type TestVars map[string]decimal.Decimal
	type TestData struct {
		input  string
		vars   TestVars
		output decimal.Decimal
	}

	data := []TestData{
		{"x+y", TestVars{"x": decimal.NewFromFloat(7.7), "y": decimal.NewFromFloat(1.2)}, decimal.NewFromFloat(8.9)},
		{"x+(-y)", TestVars{"x": decimal.NewFromFloat(100.0), "y": decimal.NewFromFloat(12.0)}, decimal.NewFromFloat(88)},
		{"x1*(x2^2)", TestVars{"x1": decimal.NewFromFloat(-100.0), "x2": decimal.NewFromFloat(7.0)}, decimal.NewFromFloat(-4900)},
		{"(доход-расход)*налог", TestVars{"доход": decimal.NewFromInt(1520), "расход": decimal.NewFromInt(840), "налог": decimal.NewFromFloat(0.87)}, decimal.NewFromFloat(591.6)},
	}

	pars := expp.NewParser()

	for _, d := range data {
		_, err := pars.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res, err := pars.Evaluate(d.vars)
		if err != nil {
			t.Error(err)
		}
		if !fuzzyEqual(res, d.output) {
			t.Error("incorrect result, need: " + d.output.String() + ", but get: " + res.String())
		}
	}
}

func TestParenthesisIsCorrect(t *testing.T) {
	type TestData struct {
		s       string
		correct bool
	}

	data := []TestData{
		{"", true},
		{"()", true},
		{")(", false},
		{"func2(600, 60, 6)", true},
		{"(func2(600, 60, 6))", true},
		{"func2(func2(700,70,7), 222, -8)", true},
		{"func2(func2(700,70,7), 222, -8", false},
		{"func2(func2(700,70,7), 222, -8))", false},
		{"(func2(func2(700,70,7), 222, -8)", false},
	}
	for i, d := range data {
		if _, cor := internal.ParenthesisIsCorrect(d.s); cor != data[i].correct {
			t.Error("incorrect result for " + strconv.Itoa(i) + " case: '" + d.s +
				"'. Need: " + strconv.FormatBool(data[i].correct) +
				", but get: " + strconv.FormatBool(cor))
		}
	}

}

func Foo1(args ...decimal.Decimal) (decimal.Decimal, error) {
	return decimal.NewFromFloat(0.1), nil
}
func TestUnaryOperatorExist(t *testing.T) {
	p := expp.NewParser()
	p.AddFunction(Foo1, "foo1")
	_, exist := internal.UnaryOperatorExist("foo1", p)
	if !exist {
		t.Error("func not found")
	}

	_, exist = internal.UnaryOperatorExist("bar", p)
	if exist {
		t.Error("func false found")
	}
}

func TestBinaryOperatorExist(t *testing.T) {
	p := expp.NewParser()
	p.AddFunction(Foo1, "foo1")
	_, exist := internal.BinaryOperatorExist("^", p)
	if !exist {
		t.Error("operator not found")
	}

	_, exist = internal.BinaryOperatorExist("~", p)
	if exist {
		t.Error("operator false found")
	}
}
