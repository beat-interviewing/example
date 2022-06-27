# Example Coding Assignment in Go

In mathematics, the Fibonacci numbers, commonly denoted _F(n)_ , form a
sequence, the Fibonacci sequence, in which each number is the sum of the two
preceding ones. 

The sequence commonly starts from the two numbers equal to 1. 

## Definition

The Fibonacci numbers may be defined by the recurrence below

_F(0) = 0, F(1) = 1_

and 

_F(n) = F(n-1) + F(n-2), for n > 1_

## Assignment

Inside [`main.go`](main.go) a stub implementation of `FibFn` is defined. Your
assignment is to implement the function so that it satisfies the requirements as
described above.

Each invocation of `FibFn` should return the next number in the sequence as
illustrated by the example below.

```go
fn := FibFn()
fn() // 1
fn() // 1
fn() // 2
fn() // 3
fn() // 5
fn() // 8
```
