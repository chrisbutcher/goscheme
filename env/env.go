package env

// import "fmt"

import "github.com/chrisbutcher/goscheme/types"

func Get(symbol types.Symbol, env *types.Environment) *types.Environment {
  if _, found := env.EnvVars[symbol]; found {
    return env
  } else {
    if env.Parent != nil {
      return Get(symbol, env.Parent)
    }
    return nil
  }
}

func Initialize(env *types.Environment) {
  env.EnvVars = make(map[types.Symbol]types.Expression)
  env.EnvVars["+"] = Add
  env.EnvVars["-"] = Subtract
  env.EnvVars["*"] = Multiply
  env.EnvVars["/"] = Divide
  env.EnvVars[">"] = GreaterThan
  env.EnvVars["<"] = LessThan
  env.EnvVars["="] = Equal
  env.EnvVars["equal?"] = Equal
  env.EnvVars["cons"] = Cons
  env.EnvVars["car"] = Car
  env.EnvVars["cdr"] = Cdr
  env.EnvVars["list"] = List
}

func Add(expression ...types.Expression) types.Expression {
  sum := expression[0].(types.Number)
  for _, num := range expression[1:] {
    sum += num.(types.Number)
  }
  return sum
}

func Subtract(expression ...types.Expression) types.Expression {
  difference := expression[0].(types.Number)
  for _, num := range expression[1:] {
    difference -= num.(types.Number)
  }
  return difference
}

func Multiply(expression ...types.Expression) types.Expression {
  product := expression[0].(types.Number)
  for _, num := range expression[1:] {
    product *= num.(types.Number)
  }
  return product
}

func Divide(expression ...types.Expression) types.Expression {
  quotient := expression[0].(types.Number)
  for _, num := range expression[1:] {
    quotient /= num.(types.Number)
  }
  return quotient
}

func GreaterThan(expression ...types.Expression) types.Expression {
  x := expression[0].(types.Number)
  y := expression[1].(types.Number)

  return types.Boolean(x > y)
}

func LessThan(expression ...types.Expression) types.Expression {
  x := expression[0].(types.Number)
  y := expression[1].(types.Number)

  return types.Boolean(x < y)
}

func Equal(expression ...types.Expression) types.Expression {
  x := expression[0].(types.Number)
  y := expression[1].(types.Number)

  return types.Boolean(x == y)
}

func Cons(expression ...types.Expression) types.Expression {
  car := expression[0]

  switch cdr := expression[1].(type) {
  case []types.Expression:
    newCdr := make([]types.Expression, 0)
    newCdr = append(newCdr, car)
    for _, exp := range cdr {
      newCdr = append(newCdr, exp)
    }
    return newCdr
  default:
    return []types.Expression{car, cdr}
  }
}

func Car(expression ...types.Expression) types.Expression {
  x := expression[0]
  car := x.([]types.Expression)[0]
  return car
}

func Cdr(expression ...types.Expression) types.Expression {
  x := expression[0]
  cdr := x.([]types.Expression)[1:]
  return cdr
}

func List(expression ...types.Expression) types.Expression {
  return expression
}
