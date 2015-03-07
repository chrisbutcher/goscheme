package main

import (
  "bufio"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
)

import (
  "github.com/chrisbutcher/goscheme/env"
  "github.com/chrisbutcher/goscheme/eval"
  "github.com/chrisbutcher/goscheme/lexer"
  "github.com/chrisbutcher/goscheme/parser"
  "github.com/chrisbutcher/goscheme/types"
  "github.com/chrisbutcher/goscheme/ui"
)

func ReadEval(input string, en types.Environment) types.Expression {
  tokens, err := lexer.Lex(input)
  if err != nil {
    fmt.Println("Lexing error: ", err)
  }

  parsed, err := parser.Parse(&tokens)
  if err != nil {
    fmt.Println("Parsing error: ", err)
  }

  evaluated, err := eval.Eval(parsed, &en)
  if err != nil {
    fmt.Println(err)
  }

  return evaluated
}

func printEvaluated(evaluated types.Expression) {
  fmt.Print(" => ", evaluated, "\n")
}

func main() {
  globals := types.Environment{}
  env.Initialize(&globals)

  flag.Parse()
  if len(flag.Args()) < 1 {
    ui.PrintBanner()
    reader := bufio.NewReader(os.Stdin)

    for {
      input, err := reader.ReadString('\n')

      if err != nil {
        fmt.Println("User input error")
      }

      if input == "exit\n" {
        return
      }

      printEvaluated(ReadEval(input, globals))
    }
  } else {
    fileBytes, err := ioutil.ReadFile(flag.Args()[0])

    if err != nil {
      fmt.Println("File input error. Exiting.")
      return
    }

    input := string(fileBytes)
    fmt.Println("> " + input)

    printEvaluated(ReadEval(input, globals))
  }
}
