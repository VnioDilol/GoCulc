package main

import (
	"fmt"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/app"

	//"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	//"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Knetic/govaluate"
)

func Cos(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Cos() requires one argument")
	}
	x, ok := args[0].(float64)
	if !ok {
		return nil, fmt.Errorf("Cos() requires a float64 argument")
	}
	return math.Cos(x), nil
}

func Sin(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Cos() requires one argument")
	}
	x, ok := args[0].(float64)
	if !ok {
		return nil, fmt.Errorf("Cos() requires a float64 argument")
	}
	return math.Sin(x), nil
}

func Tan(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Cos() requires one argument")
	}
	x, ok := args[0].(float64)
	if !ok {
		return nil, fmt.Errorf("Cos() requires a float64 argument")
	}
	return math.Tan(x), nil
}

func Log(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Cos() requires one argument")
	}
	x, ok := args[0].(float64)
	if !ok {
		return nil, fmt.Errorf("Cos() requires a float64 argument")
	}
	return math.Log(x), nil
}

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Калькулятор")

	inputLabel := widget.NewLabel("")
	resultLabel := widget.NewLabel("")
	resultLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}

	grid := container.NewGridWithColumns(4,
		widget.NewButton("1", func() { inputLabel.SetText(inputLabel.Text + "1") }),
		widget.NewButton("2", func() { inputLabel.SetText(inputLabel.Text + "2") }),
		widget.NewButton("3", func() { inputLabel.SetText(inputLabel.Text + "3") }),
		widget.NewButton("+", func() { inputLabel.SetText(inputLabel.Text + "+") }),

		widget.NewButton("4", func() { inputLabel.SetText(inputLabel.Text + "4") }),
		widget.NewButton("5", func() { inputLabel.SetText(inputLabel.Text + "5") }),
		widget.NewButton("6", func() { inputLabel.SetText(inputLabel.Text + "6") }),
		widget.NewButton("-", func() { inputLabel.SetText(inputLabel.Text + "-") }),

		widget.NewButton("7", func() { inputLabel.SetText(inputLabel.Text + "7") }),
		widget.NewButton("8", func() { inputLabel.SetText(inputLabel.Text + "8") }),
		widget.NewButton("9", func() { inputLabel.SetText(inputLabel.Text + "9") }),
		widget.NewButton("*", func() { inputLabel.SetText(inputLabel.Text + "*") }),

		widget.NewButton("C", func() { inputLabel.SetText(""); resultLabel.SetText("") }),
		widget.NewButton("0", func() { inputLabel.SetText(inputLabel.Text + "0") }),
		widget.NewButton("=", func() {
			expr, err := govaluate.NewEvaluableExpressionWithFunctions(inputLabel.Text, map[string]govaluate.ExpressionFunction{
				"Sin": Sin,
				"Cos": Cos,
				"Tg":  Tan,
				"Ln":  Log,
			})
			if err == nil {
				result, err := expr.Evaluate(nil)
				if err == nil {
					resultLabel.SetText(fmt.Sprintf("= %v", result))
				} else {
					resultLabel.SetText(fmt.Sprintf("Ошибка: %v", err))
				}
			} else {
				resultLabel.SetText(fmt.Sprintf("Ошибка: %v", err))
			}
		}),
		widget.NewButton("/", func() { inputLabel.SetText(inputLabel.Text + "/") }),
	)

	trigonometryBox := container.NewGridWithColumns(7,
		widget.NewButton("sin", func() { inputLabel.SetText(inputLabel.Text + "Sin(") }),
		widget.NewButton("cos", func() { inputLabel.SetText(inputLabel.Text + "Cos(") }),
		widget.NewButton("tan", func() { inputLabel.SetText(inputLabel.Text + "Tan(") }),
		widget.NewButton("ln", func() { inputLabel.SetText(inputLabel.Text + "Ln(") }),
		widget.NewButton("pi", func() { inputLabel.SetText(inputLabel.Text + "3.14") }),
		widget.NewButton("(", func() { inputLabel.SetText(inputLabel.Text + "(") }),
		widget.NewButton(")", func() { inputLabel.SetText(inputLabel.Text + ")") }),
	)

	content := container.NewVBox(
		widget.NewLabel("Калькулятор"),
		widget.NewCard("", "", grid),
		widget.NewCard("", "", trigonometryBox),
		inputLabel,
		resultLabel,
	)
	myWin.SetContent(content)
	myWin.ShowAndRun()
}
