# goscheme
An experimental Scheme interpreter and REPL in Go.

[![GoDoc](https://godoc.org/github.com/chrisbutcher/goscheme?status.svg)](https://godoc.org/github.com/chrisbutcher/goscheme)

## Running
```bash
cd goscheme
go install
goscheme # for REPL
goscheme sample.scm # to run a Scheme script
```

## Examples
```scheme
> (+ 3 4)
 => 7

> (car (cdr (cdr (list 1 2 3))))
 => 3

> (begin 
    (define square (lambda (x) (* x x)))
    (define sum-of-squares (lambda (a b) (+ (square a) (square b))))
    (sum-of-squares 3 4))
 => 25

> (define y 12)
 => 12
> (set! y 41)
 => 41
> ((lambda (x) (+ x y)) 1)
 => 42

> (((lambda (x) (x x))
    (lambda (fact-gen)
      (lambda (n)
        (if (= 1 n) 1 (* n ((fact-gen fact-gen) (- n 1))))
      )
    )
  ) 100)
 => 9.33262154439441e+157

> (begin
    (define lessthanten (lambda (x) (if (< x 10) #t #f)))
    (lessthanten 9))
 => #t

```
