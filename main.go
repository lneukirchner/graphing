package main

import (
	"fmt"
	"math"
	"time"
)

const precision int = 4096

//Elementary function definitions
func factorial(x float64) float64 {
	var i float64 = x
	if x > 1 {
		i *= factorial(x - 1)
	} else {
		i = 1
	}
	return i
}

func sin(x float64) float64 {
	var result float64
	var i int
	for i = 0; i < 20; i++ {
		addend := ((math.Pow(-1, float64(i))) / (factorial(2*float64(i) + 1))) * math.Pow(x, float64(2*i+1))
		fmt.Println(addend)
		result = result + addend
	}
	return result
}

//func cos(x float64) float64 {

//}

//Working function definition
func f(x float64) float64 {
	return (x*x + x + 1)
}

//Function manipulation definitions
func derivativeApproximation(x float64) float64 {
	firstPoint := [2]float64{x, f(x)}
	secondX := x + 1/float64(precision)
	secondPoint := [2]float64{secondX, f(secondX)}
	derivative := (secondPoint[1] - firstPoint[1]) / (secondPoint[0] - firstPoint[0])
	return derivative
}

func lhIntegrate(leftEnd, rightEnd float64) float64 {
	delta := (rightEnd - leftEnd) / float64(precision)
	var j int
	var approx float64
	for j = 0; j < precision; j++ {
		approx += f((float64(j)*delta)+leftEnd) * delta
	}
	return approx
}

func trapezoidArea(b1, b2, h float64) float64 {
	return ((b1 + b2) / 2) * h
}

func rhIntegrate(leftEnd, rightEnd float64) float64 {
	delta := (rightEnd - leftEnd) / float64(precision)
	var j int
	var approx float64
	for j = 1; j <= precision; j++ {
		approx += f((float64(j)*delta)+leftEnd) * delta
	}
	return approx
}

func trapezoidIntegrate(leftEnd, rightEnd float64) float64 {
	delta := (rightEnd - leftEnd) / float64(precision)
	var j int
	var approx float64
	for j = 0; j < precision; j++ {
		leftSide := f(float64(j)*delta + leftEnd)
		rightSide := f(float64(j+1)*delta + leftEnd)
		base := delta
		approx += trapezoidArea(leftSide, rightSide, base)
	}
	return approx
}

func main() {
	//Stopwatch code for execution time
	start := time.Now()
	defer fmt.Println("Program took", time.Since(start))

	//code to check
	answer := sin(math.Pi / 6)
	fmt.Println(answer)
}
