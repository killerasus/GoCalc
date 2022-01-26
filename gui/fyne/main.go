package main

import (
	"calculator/calculator"
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func calcFormat(a float64) string {
	return strconv.FormatFloat(a, 'f', 4, 64)
}

func main() {
	var calc calculator.Calculator

	calcApp := app.New()
	window := calcApp.NewWindow("Go RPN Calculator")
	window.SetTitle("Go RPN Calculator")
	icon, err := fyne.LoadResourceFromPath("./resources/calculator.png") // Expects being run from root folder
	if err != nil {
		fyne.LogError("Couldn't load app icon.", err)
	} else {
		window.SetIcon(icon)
	}

	memory := widget.NewLabel("Memory:")
	stack := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{})

	updateMemory := func() {
		stack.SetText("")
		for i := 0; i < calc.Size(); i++ {
			v, _ := calc.Peeki(i)
			stack.SetText(fmt.Sprint(stack.Text, fmt.Sprint(v), "\n"))
		}
	}

	display := widget.NewLabelWithStyle("0", fyne.TextAlignTrailing, fyne.TextStyle{})

	addNumber := func(v float64) {
		if n, err := strconv.ParseFloat(display.Text, 64); err == nil {
			if n == 0.0 {
				display.SetText(calcFormat(v))
			} else {
				display.SetText(fmt.Sprint(n, calcFormat(v)))
			}
		} else {
			display.SetText((calcFormat(v)))
		}
	}

	keypad := container.New(
		layout.NewGridLayout(4),
		widget.NewButton("7", func() { addNumber(7) }),
		widget.NewButton("8", func() { addNumber(8) }),
		widget.NewButton("9", func() { addNumber(9) }),
		widget.NewButton("+", func() {
			if _, ok := calc.Add(); !ok {
				display.SetText("Error")
			}
			updateMemory()
		}),
		widget.NewButton("4", func() { addNumber(4) }),
		widget.NewButton("5", func() { addNumber(5) }),
		widget.NewButton("6", func() { addNumber(6) }),
		widget.NewButton("-", func() {
			if _, ok := calc.Subtract(); !ok {
				display.SetText("Error")
			}
			updateMemory()
		}),
		widget.NewButton("1", func() { addNumber(1) }),
		widget.NewButton("2", func() { addNumber(2) }),
		widget.NewButton("3", func() { addNumber(3) }),
		widget.NewButton("*", func() {
			if _, ok := calc.Multiplication(); !ok {
				display.SetText("Error")
			}
			updateMemory()
		}),
		widget.NewButton("0", func() {
			if v, err := strconv.ParseFloat(display.Text, 64); err != nil || v == 0.0 {
				display.SetText("0")
			} else {
				addNumber(0)
			}
		}),
		widget.NewButton(".", func() {
			// TODO
		}),
		widget.NewButton("Enter", func() {
			v, err := strconv.ParseFloat(display.Text, 64)
			if err == nil {
				calc.Push(v)
			}
			display.SetText("0")
			updateMemory()
		}),
		widget.NewButton("/", func() {
			if v, ok := calc.Division(); !ok {
				if v == math.Inf(0) {
					display.SetText("Division by zero")
				} else {
					display.SetText("Error")
				}
			}
			updateMemory()
		}),
	)

	input := container.NewVBox(display, keypad)
	memoryBox := container.NewVBox(memory, stack)
	box := container.NewHBox(memoryBox, input)

	window.SetContent(box)

	window.Show()
	calcApp.Run()
}
