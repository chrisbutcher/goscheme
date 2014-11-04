package main

import "fmt"

type Expr []T

func (x Expr) ToAtomSlice() []Atom {
  atoms := make([]Atom, 0)

  for _, exp := range x {
    atoms = append(atoms, exp.(Atom))
  }

  return atoms
}

func (x Expr) Eval(env Environment) T {
  if len(x) == 1 {
    switch x[0].(type) {
    case Expr: // Nested exression, further evaluation is needed
      return x[0].(Expr).Eval(env)
    case Atom: // Found Atom
      atom := x[0].(Atom)

      if atom.typ == atomSymbol {
        res, found := env.get(atom.val.(string))

        if found {
          return res
        } else {
          return atom
        }

      } else {
        return atom
      }
    default:
      fmt.Printf("Unknown.")
    }
    return x
  } else {
    evaluated_expressions := make(Expr, 0)

    if exprIsLambda(x[0]) {
      x[0] = populateLambda(x)
      return x[0]
    } else if exprIsBoolean(x[0]) {
      return x[0].(Atom).val
    } else {
      for index, _ := range x {
        evaluated_expressions = append(evaluated_expressions, x[index:index+1].Eval(env))
      }

      operands := evaluated_expressions[1:].ToAtomSlice()
      fnAtom := evaluated_expressions[0].(Atom)

      switch fnAtom.typ {
      case atomLambda:
        return fnAtom.lambdaFn.(Expr).Eval(New(fnAtom.lambdaArgs, operands, env))
      case atomBuiltin:
        return fnAtom.lambdaFn.(func(input []Atom, env Environment) Atom)(operands, env)
      default:
        fmt.Println("Error: failed to evaluate input")
        return fnAtom
      }
    }
  }
}
