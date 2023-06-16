package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}

type TerminalExpression struct {
	str string
}

func (exp TerminalExpression) Interpret() int {
	number, _ := strconv.Atoi(exp.str)

	return number
}

type AddExpression struct {
	exp1 Expression
	exp2 Expression
}

func (e AddExpression) Interpret() int {
	return e.exp1.Interpret() + e.exp2.Interpret()
}

type SubExpression struct {
	exp1 Expression
	exp2 Expression
}

func (e SubExpression) Interpret() int {
	return e.exp1.Interpret() - e.exp2.Interpret()
}

type Calculator struct{}

func (Calculator) parse(str string) Expression {
	strs := strings.Split(str, " ")

	if len(strs) == 1 {
		return TerminalExpression{str: str}
	}

	var numbStrings []string
	var opsStrings []string

	for _, s := range strs {
		switch s {
		case "+":
			fallthrough
		case "-":
			opsStrings = append(opsStrings, s)
		default:
			numbStrings = append(numbStrings, s)
		}
	}

	var currentExp Expression

	for i := 0; i < len(opsStrings); i++ {
		if i == 0 {
			currentExp = TerminalExpression{str: numbStrings[0]}
		}

		if opsStrings[i] == "+" {
			currentExp = AddExpression{
				exp1: currentExp,
				exp2: TerminalExpression{str: numbStrings[i+1]},
			}
		}

		if opsStrings[i] == "-" {
			currentExp = SubExpression{
				exp1: currentExp,
				exp2: TerminalExpression{str: numbStrings[i+1]},
			}
		}
	}

	return currentExp
}

func (cal Calculator) Result(str string) int {
	return cal.parse(str).Interpret()
}

func main() {
	calculator := Calculator{}

	fmt.Println("3 =", calculator.Result("3"))
	fmt.Println("3 + 7 + 2 =", calculator.Result("3 + 7 + 2"))
	fmt.Println("5 - 1 + 8 =", calculator.Result("5 - 1 + 8"))
	fmt.Println("9 - 2 - 3 =", calculator.Result("9 - 2 - 3"))
	fmt.Println("8 + 1 - 9 =", calculator.Result("8 + 1 - 9"))
}
