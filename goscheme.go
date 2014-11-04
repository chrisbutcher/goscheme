package main

import (
  "bufio"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
)

type T interface{}

func interpretInput(input string, env Environment) {
  tokens := Tokenize(input)
  parseTree, _ := Parenthesize(tokens)
  result := parseTree.Eval(env).(Atom)

  fmt.Print(" => ")

  switch result.typ {
  case atomFloat:
    fmt.Printf("%f\n", result.valNum)
  case atomQuote:
    fmt.Printf("%v\n", result.val)
  case atomBoolean:
    fmt.Printf("%v\n", result.val)
  default:
    fmt.Printf("%v\n", result.val)
  }
}

func main() {
  printBanner()
  reader := bufio.NewReader(os.Stdin)

  env := Environment{}
  env.initialize()

  flag.Parse()

  if len(flag.Args()) < 1 {
    for {
      fmt.Print("> ")
      input, err := reader.ReadString('\n')

      if err != nil {
        fmt.Println("User input error. Exiting.")
        continue
      }

      if input == "exit\n" {
        return
      }

      if !validParens(input) {
        fmt.Println("User input error. The number of opening and closing parentheses do not match.")
        continue
      }

      input = addOutsideParamsIfMissing(input)

      if input != "\n" {
        interpretInput(input, env)
      }
    }
  } else {
    fileBytes, err := ioutil.ReadFile(flag.Args()[0])
    if err != nil {
      fmt.Println("File input error. Exiting.")
      return
    }

    fileInputString := string(fileBytes)
    fmt.Println("> " + fileInputString)

    interpretInput(fileInputString, env)
  }
}
