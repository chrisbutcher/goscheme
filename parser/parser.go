package parser

import (
  "errors"
  "strconv"
)

import "github.com/chrisbutcher/goscheme/types"

func Parse(tokens *[]string) (types.Expression, error) {
  token := (*tokens)[0]
  *tokens = (*tokens)[1:]

  if token == "(" {
    list := make([]types.Expression, 0)

    for (*tokens)[0] != ")" {
      if parsed, err := Parse(tokens); err == nil {
        list = append(list, parsed)
      } else {
        return nil, err
      }
    }
    *tokens = (*tokens)[1:]
    return list, nil
  } else if token == ")" {
    return nil, errors.New("Unexpected )")
  } else {
    if token == "#t" {
      return types.Boolean(true), nil
    } else if token == "#f" {
      return types.Boolean(false), nil
    } else if number, err := strconv.ParseFloat(token, 64); err == nil {
      return types.Number(number), nil
    } else {
      return types.Symbol(token), nil
    }
  }
}
