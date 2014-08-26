package main

import "fmt"

func Eval(x Sexpr, env Environment) T {
	if len(x) == 1 {
		switch x[0].(type) {
		case Sexpr: // Nested exression, further evaluation is needed
			return Eval(x[0].(Sexpr), env)
		case Atom: // Found Atom
			a := x[0].(Atom)

			if a.typ == atomSymbol {
				res, found := env.get(a.val.(string))

				if found {
					return res
				} else {
					return a
				}

			} else {
				return a
			}
		default:
			fmt.Printf("Unknown.")
		}
		return x
	} else {
		evaled_expressions := make(Sexpr, 0)

		if exprIsLambda(x[0]) {
			x[0] = populateLambda(x)
			return x[0]
		} else if exprIsBoolean(x[0]) {
			return x[0].(Atom).val
		} else {
			for index, _ := range x {
				evaled_expressions = append(evaled_expressions, Eval(x[index:index+1], env))
			}

			a := make([]Atom, 0)

			for _, exp := range evaled_expressions[1:] {
				a = append(a, exp.(Atom))
			}

			fnAtom := evaled_expressions[0].(Atom)

			if fnAtom.typ == atomLambda {
				lambdaFn := make(Sexpr, 0)

				for _, lambdaAtom := range fnAtom.lambdaFn {
					lambdaFn = append(lambdaFn, lambdaAtom)
				}

				return Eval(lambdaFn, New(fnAtom.lambdaArgs, a, env))
			} else {
				fn := builtins[fnAtom.val.(string)]

				if fn != nil {
					result := fn(a, env)
					return result
				} else {
					return fnAtom // Need to return error atom instead
				}
			}
		}
	}
}
