# goscheme
An experimental Scheme interpreter and REPL in Go.

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

> (car (cdr (cdr (list 1 2 3))))
 => 3

> (define timestwo (lambda (x) (* x 2)))
 => timestwo
> (timestwo 50)
 => 100

> (define y 12)
 => 12
> (set! y 41)
 => 41
> ((lambda (x) (+ x y)) 1)
 => 42

> (begin (define lessthanten (lambda (x) (if (< x 10) #t #f))) (lessthanten 9))
 => #t

```
