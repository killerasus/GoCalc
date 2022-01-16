package main

import (
	"bufio"
	"calculator/calculator"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintHelp() {
	fmt.Printf("\t<number>\tPushes <number> into stack\n")
	fmt.Printf("\t+\t\tAdds two numbers on the stack\n")
	fmt.Printf("\t-\t\tSubtracts two numbers on the stack or applies unary minus\n")
	fmt.Printf("\thelp\t\tPrints this help\n")
	fmt.Printf("\tCtrl+Z\t\tEnds the calculator\n")
}

func main() {

	help := flag.Bool("h", false, "This help")
	flag.Parse()

	if *help {
		fmt.Println("Go RPN calculator:")
		PrintHelp()
	}

	var calc calculator.Calculator
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		input = strings.ToLower(scanner.Text())

		switch input {
		case "+":
			if value, ok := calc.Add(); ok {
				fmt.Printf("\t%f\n", value)
			} else {
				fmt.Println("Error: Stack size is less than 2:", calc.Size())
			}
		case "-":
			if value, ok := calc.Subtract(); ok {
				fmt.Printf("\t%f\n", value)
			} else {
				fmt.Println("Error: Stack size is less than 1:", calc.Size())
			}
		case "help":
			PrintHelp()
		default:
			value, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				calc.Push(value)
				v, _ := calc.Peek()
				fmt.Println("Pushed to stack:", v, "Stack size:", calc.Size())
			}
		}
		fmt.Print("> ")
	}
}