package parser

import (
	"errors"
	"sort"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

const float64EqualityThreshold = 1e-9

func TestGetVarList(t *testing.T) {
	allElemsIsUnique := func(arr []string) error {
		for i, v := range arr {
			for j := i + 1; j < len(arr); j++ {
				if v == arr[j] {
					return errors.New("non unique var: " + v)
				}
			}
		}
		return nil
	}
	isEqual := func(arr1, arr2 []string) error {
		if len(arr1) != len(arr2) {
			return errors.New("different arrays size. " +
				"len(arr1): " + strconv.Itoa(len(arr1)) +
				", len(arr2): " + strconv.Itoa(len(arr2)))
		}
		sort.Strings(arr1)
		sort.Strings(arr2)
		for i, v := range arr1 {
			if v != arr2[i] {
				return errors.New("different arrays elements: " + v + ", " + arr2[i])
			}
		}
		return nil
	}

	type TestData struct {
		input  string
		output []string
	}

	data := []TestData{
		{"", []string{}},
		{"x", []string{"x"}},
		{"x*(sqrt(y)+1)", []string{"x", "y"}},
		{"(доход-расход)*налог", []string{"доход", "расход", "налог"}},
	}

	parser := NewParser()

	for _, d := range data {
		exp, err := parser.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res := GetVarList(exp)
		err = allElemsIsUnique(res)
		if err != nil {
			t.Error(err)
		}
		err = isEqual(res, d.output)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewParser(t *testing.T) {
	func1 := func(args ...decimal.Decimal) (decimal.Decimal, error) {
		return args[0].Add(decimal.NewFromFloat(100)), nil
	}
	func2 := func(args ...decimal.Decimal) (decimal.Decimal, error) {
		return args[0].Add(decimal.NewFromFloat(200)), nil
	}
	parser1 := NewParser()
	parser2 := NewParser()
	parser1.AddFunction(func1, "f1")
	parser1.AddFunction(func2, "f2")

	parser2.AddFunction(func1, "f2")
	parser2.AddFunction(func2, "f1")

	data := []string{"f1(1)", "f2(2)"}

	parser1.Parse(data[0])
	res, err := parser1.Evaluate(map[string]decimal.Decimal{"a": decimal.NewFromFloat(1)})
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(101)) {
		t.Error("incorrect parser1 result, need: 101.0, but get: " + res.String())
	}

	parser1.Parse(data[1])
	res, err = parser1.Evaluate(map[string]decimal.Decimal{})
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(202)) {
		t.Error("incorrect parser1 result, need: 202.0, but get: " + res.String())
	}

	parser2.Parse(data[0])
	res, err = parser2.Evaluate(map[string]decimal.Decimal{})
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(201)) {
		t.Error("incorrect parser2 result, need: 201.0, but get: " + res.String())
	}

	parser2.Parse(data[1])
	res, err = parser2.Evaluate(map[string]decimal.Decimal{})
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(102)) {
		t.Error("incorrect parser2 result, need: 102.0, but get: " + res.String())
	}
}

func TestNewParser2(t *testing.T) {
	parser := NewParser()
	type TestData struct {
		input  string
		output decimal.Decimal
	}
	data := []TestData{
		{"15+20", decimal.NewFromInt(35)},
		{"2^3-10", decimal.NewFromInt(-2)},
		{"sqrt(14+(4^(0.5)))", decimal.NewFromInt(4)},
	}

	for _, d := range data {
		_, err := parser.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res, err := parser.Evaluate(map[string]decimal.Decimal{})
		if err != nil {
			t.Error(err)
		}
		if !res.Equal(d.output) {
			t.Error("incorrect result, need: " + d.output.String() + ", but get: " + res.String())
		}
	}
}

func TestParse(t *testing.T) {
	type TestData struct {
		input  string
		output decimal.Decimal
	}

	data := []TestData{
		{"", decimal.Zero},
		{"10+50+5", decimal.NewFromInt(65)},
		{"2*2+2", decimal.NewFromInt(6)},
		{"2*(2+2)", decimal.NewFromInt(8)},
		{"3^2", decimal.NewFromInt(9)},
		{"2*2+3", decimal.NewFromInt(7)},
		{"sqrt(3^2+(2*2+3))", decimal.NewFromInt(4)},
		{"100+sqrt(3^2+(2*2+3))", decimal.NewFromInt(104)},
		{"2*-1", decimal.NewFromInt(-2)},
	}
	parser := NewParser()
	for _, d := range data {
		_, err := parser.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res, err := parser.Evaluate(map[string]decimal.Decimal{})
		if err != nil {
			t.Error(err)
		}
		if !res.Equal(d.output) {
			t.Error("incorrect result, need: " + d.output.String() + ", but get: " + res.String())
		}
	}
}

func TestParserString(t *testing.T) {
	p := NewParser()
	p.Parse("")
	if p.String() != "0" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("1 + a")
	if p.String() != "( + 1 a )" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("2^(sqrt(14+2))")
	if p.String() != "( ^ 2 ( sqrt ( ( + 14 2 ) ) ) )" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("abs(- 2)")
	if p.String() != "( abs ( ( - 2 ) ) )" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("- 4")
	if p.String() != "( - 4 )" {
		t.Error("incorrect string conversion = " + p.String())
	}
}

func TestParse2(t *testing.T) {
	p := NewParser()
	exp, err := p.Parse("2 * (x+y")
	if exp != nil || err == nil {
		t.Error("incorrect error handling")
	}
	exp, err = p.Parse("Foo(x+y)")
	if exp != nil || err == nil {
		t.Error("incorrect error handling")
	}
}

func TestParseStr(t *testing.T) {
	//p := Parser{}
	//i, err := p.parseStr([]rune("asf"))
	// TODO: finish
}

func TestParseFunc(t *testing.T) {
	//p := Parser{}
	//_, isFunc, err := p.parseFunc([]rune("foo(a+b)"))
	// TODO: finish

}
