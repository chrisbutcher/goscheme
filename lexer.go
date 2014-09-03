package main

import "strings"

// Returns nested token array, and mutated (popped from front) tokens for recursion
func Parenthesize(tokens Expr) (Expr, Expr) {
	var t T = tokens[0] // The current token

	if t.(string) == "(" {
		ast := make(Expr, 0) // Abstract syntax tree, nested strings and slices
		for {
			tokens = tokens[1:]
			next_token_or_expression, new_tokens := Parenthesize(tokens)

			tokens = new_tokens
			next_token := next_token_or_expression[0]

			if next_token == ")" {
				break
			}

			if len(next_token_or_expression) == 1 {
				ast = append(ast, createAtom(next_token))
			} else {
				ast = append(ast, next_token_or_expression)
			}
		}
		return ast, tokens

	} else {
		// Must return ast as Expr, even if it is a single token
		ast := make(Expr, 0)
		ast = append(ast, t)

		return ast, tokens
	}
}

func Tokenize(s string) Expr {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)

	split_tokens := strings.Split(s, " ")

	tokens := make(Expr, 0)

	for _, t := range split_tokens {
		if t != "" && t != "\n" {
			tokens = append(tokens, t)
		}
	}

	return tokens
}
