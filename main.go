package main

import "fmt"

func f(x float64) float64 {
	return (x*x + x + 1)
}

func derivativeApproximation(x float64, precision float64) float64 {
	firstPoint := [2]float64{x, f(x)}
	secondX := x + 1/precision
	secondPoint := [2]float64{secondX, f(secondX)}
	derivative := (secondPoint[1] - firstPoint[1]) / (secondPoint[0] - firstPoint[0])
	return derivative
}

func lhIntegrate(leftEnd, rightEnd float64, precision int) float64 {
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

func rhIntegrate(leftEnd, rightEnd float64, precision int) float64 {
	delta := (rightEnd - leftEnd) / float64(precision)
	var j int
	var approx float64
	for j = 1; j <= precision; j++ {
		approx += f((float64(j)*delta)+leftEnd) * delta
	}
	return approx
}

func trapezoidIntegrate(leftEnd, rightEnd float64, precision int) float64 {
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
	var leftEnd, rightEnd float64 = 1.0, 3.0
	var precision int = 4096 //precision either represents the amount of sectors in an integral, or 1/dx
	leftHandApprox := lhIntegrate(leftEnd, rightEnd, precision)
	rightHandApprox := rhIntegrate(leftEnd, rightEnd, precision)
	trapApprox := trapezoidIntegrate(leftEnd, rightEnd, precision)
	fmt.Println("Left hand approximation is", leftHandApprox, "and the right hand approximation is", rightHandApprox)
	average := (leftHandApprox + rightHandApprox) / 2
	fmt.Printf("The average is %.2f\n", average)
	fmt.Printf("The trapezoid approximation is %.2f\n", trapApprox)
	d1 := derivativeApproximation(leftEnd, float64(precision))
	d2 := derivativeApproximation(rightEnd, float64(precision))
	fmt.Printf("The derivative at the endpoints is %.2f at x = %f and %.2f at x = %f\n", d1, leftEnd, d2, rightEnd)
}
