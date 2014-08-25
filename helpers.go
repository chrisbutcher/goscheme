package main

import "fmt"
import "strings"

func printBanner() {
	fmt.Println("===============================")
	fmt.Println("       GoScheme REPL v0.1")
	fmt.Println("(Ctrl+C or type 'exit' to Quit)")
	fmt.Println("===============================")
}

func validParens(input string) bool {
	a := matchingParensCount(input)
	b := matchingParens(input)

	return a && b
}

func matchingParensCount(input string) bool {
	num_opening := strings.Count(input, "(")
	num_closing := strings.Count(input, ")")

	return num_opening == num_closing
}

func matchingParens(input string) bool {
	stack := 0

	for _, char := range input {
		if char == '(' {
			stack += 1
		} else if char == ')' {
			stack -= 1
		}
	}

	return stack == 0
}

func addOutsideParamsIfMissing(input string) string {
	input = strings.Trim(input, " \n")

	if !strings.HasPrefix(input, "(") && !strings.HasSuffix(input, ")\n") {
		input = "(" + input + ")"
	}

	return input
}
