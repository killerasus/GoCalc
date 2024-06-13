# Go RPN Calculator

An over-engineered RPN calculator built to practice unit testing, error reporting, GUI development, and CLI in Go.

[![Build Status](https://app.travis-ci.com/killerasus/GoCalc.svg?branch=master)](https://app.travis-ci.com/github/killerasus/GoCalc)

## CLI

To run the command line interface of the Go RPN Calculator, go to folder `cli` and run `go run cli.go` or `go build`.

CLI interface supports flags. See implemented flags and usage with `-h` or `-help`.

For example: CLI supports expression evaluation. The result is stacked.

![Using CLI for evaluating expressions](img/cli.png "Example using -eval option to evaluate RPN expression")

## GUI

### Fyne

A GUI version of the calculator made with the free and open source [fyne](https://fyne.io/) toolkit.

![Calculator GUI in Fyne](img/fynegui.png "GUI implemented using the Fyne toolkit")

Compile with `go build -o gorpn.exe gui/fyne/main.go`. Remember that you will have to provide a GNU C++ compiler (in Windows you can use [MSYS2](https://www.msys2.org/) to provide GCC).

Calculator icons created by [Freepik - Flaticon](https://www.flaticon.com/free-icons/calculator).
## Testing

Run tests with `go test ./stack ./calculator ./operations -v` command in the root directory.

## Class Diagram

```mermaid
classDiagram

    IStack <|-- Stack
    Calculator *-- Stack
    Calculator --> Operations

    class IStack {
        <<interface>>
        +Push(float64)
        +Pop() float64
        +Peek() float64
        +Peeki(int i) (float64, bool)
        +Size() int
        +IsEmpty() bool
    }

    class Stack  {
        -float64[]
    }

    class Calculator {
        -Stack stack
        +Size() int
        +Push(float64 a) void
        +Peek() (float64, bool)
        +Peeki(i int) (float64, bool)
        +Pop() (float64, bool)
        +Add() (float64, bool)
        +Subtract() (float64, bool)
        +Multiplication() (float64, bool)
        +Division() (float64, bool)
        +Evaluate(string e) (v float64, ok bool)
    }

    class Operations {
        <<namespace>>
        +Add(float64 a, float64 b) float64
        +Subtract(float64 a, float64 b) float64
        +Multiply(float64 a, float64 b) float64
        +Divide(float64 a, float64 b) float64
        +UnaryNegative(float64 a) float64
        +Root(float64 a, float64 root) float64
        +SquareRoot(float64 a) float64
    }
```
