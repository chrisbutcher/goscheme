package eval

import (
  "errors"
  "fmt"
)

import (
  "github.com/chrisbutcher/goscheme/env"
  "github.com/chrisbutcher/goscheme/types"
)

func Apply(operator types.Expression, operands []types.Expression) (types.Expression, error) {
  switch op := operator.(type) {
  case types.Lambda:
    en := &types.Environment{EnvVars: make(types.EnvVars), Parent: op.Env}

    switch args := op.Args.(type) {
    case []types.Expression:
      for i, arg := range args {
        en.EnvVars[arg.(types.Symbol)] = operands[i]
      }
    default:
      en.EnvVars[args.(types.Symbol)] = args
    }

    if evaluated, err := Eval(op.Body, en); err == nil {
      return evaluated, nil
    } else {
      return nil, err
    }
  case func(...types.Expression) types.Expression:
    return op(operands...), nil
  default:
    fmt.Println("Apply error", op)
  }
  return nil, nil
}

func Eval(expression types.Expression, en *types.Environment) (types.Expression, error) {
  switch exp := expression.(type) {
  case types.Symbol:
    found_env := env.Get(exp, en)
    if found_env != nil {
      return found_env.EnvVars[exp], nil
    } else {
      err_msg := fmt.Sprint("Eval error. Symbol not found: ", exp)
      return nil, errors.New(err_msg)
    }
  case types.Number:
    return exp, nil
  case []types.Expression:
    switch car, _ := exp[0].(types.Symbol); car {
    case "begin":
      for _, e := range exp[1:] {
        results := make([]types.Expression, 0)

        if evaluated, err := Eval(e, en); err == nil {
          fmt.Println(evaluated)
          results = append(results, evaluated)
        } else {
          return nil, nil
        }

        return results[len(results)-1], nil
      }
    case "if":
      if test, err := Eval(exp[1], en); err == nil {
        conseq := exp[2]
        alt := exp[3]

        if test.(bool) {
          return Eval(conseq, en)
        } else {
          return Eval(alt, en)
        }
      } else {
        return nil, err
      }
    case "lambda":
      return types.Lambda{exp[1], exp[2], en}, nil
    case "quote":
      return exp[1], nil
    case "define":
      if evaluated, err := Eval(exp[2], en); err == nil {
        en.EnvVars[exp[1].(types.Symbol)] = evaluated
        return exp[1].(types.Symbol), nil
      } else {
        return nil, err
      }
    default:
      operands := exp[1:]
      evaluated := make([]types.Expression, 0)
      for _, operand := range operands {
        if result, err := Eval(operand, en); err == nil {
          evaluated = append(evaluated, result)
        } else {
          return nil, err
        }
      }

      if fn, err := Eval(exp[0], en); err == nil {
        if result, err := Apply(fn, evaluated); err == nil {
          return result, nil
        } else {
          return nil, err
        }
      } else {
        return nil, err
      }
    }
  default:
    err_msg := "Eval error: " + exp.(string)
    return nil, errors.New(err_msg)
  }

  return nil, errors.New("Unexpected behavior")
}
