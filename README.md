# goscheme
An experimental Scheme interpreter and REPL in Go.

## Latest
See the [experimental](https://github.com/chrisbutcher/goscheme/tree/experimental) branch for a cleaner re-design that is underway.

## Running
```bash
cd goscheme
go install
goscheme # for REPL
goscheme sample.scm # to run a Scheme script
```

## Examples
```clojurescript
> (+ 3 4)
 => 7

> (define timestwo (lambda (x) (* x 2)))
 => lambda
> (timestwo 50)
 => 100

> (car (cdr (cdr (quote 1 2 3))))
 => 3

> ((lambda (x) (+ x 5)) ((lambda (y) (+ y 1)) 1))
 => 7

> (define y 41)
 => 41
> ((lambda (x) (+ x y)) 1)
 => 42

> (define lessthanten (lambda (x) (if (< x 10) true false)))
 => lambda
> (lessthanten 9)
 => true

```
