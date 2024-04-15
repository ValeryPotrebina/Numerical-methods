package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return -0.1027*math.Pow(x, 3) + 1.2349*math.Pow(x, 2) - 0.3256*x + 1.3851
}

//−0.1027x3+1.2349x2−0.3256x+1.3851

func main() {
	x := [...]float64{
		1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5,
	}
	y := [...]float64{
		2.35, 2.74, 5.37, 6.96, 8.52, 10.52, 13.41, 15.93, 17.61,
	}
	// Из учебника формулы
	// ср арифметическое
	xAr := (x[0] + x[8]) / 2
	// ср геометрическое
	xGeom := math.Sqrt(x[0] * x[8])
	// ср гармоническое
	xHarm := 2 / (1/x[0] + 1/x[8])

	yAr := (y[0] + y[8]) / 2
	yGeom := math.Sqrt(y[0] * y[8])
	yHarm := 2 / (1/y[0] + 1/y[8])

	zAr := f(xAr)
	zGeom := f(xGeom)
	zHarm := f(xHarm)

	fmt.Println("X : ", xAr, xGeom, xHarm)
	fmt.Println("Y : ", yAr, yGeom, yHarm)
	fmt.Println("Z : ", zAr, zGeom, zHarm)
	deltas := [...]float64{
		math.Abs(zAr - yAr),
		math.Abs(zGeom - yGeom),
		math.Abs(zAr - yGeom),
		math.Abs(zGeom - yAr),
		math.Abs(zHarm - yAr),
		math.Abs(zAr - yHarm),
		math.Abs(zHarm - yHarm),
		math.Abs(zHarm - yGeom),
		math.Abs(zGeom - yHarm),
	}
	minDelt := deltas[0]
	minI := 0

	for i, el := range deltas {
		if el < minDelt {
			minDelt = el
			minI = i
		}
	}
	fmt.Println(" Deltas : ", deltas)
	fmt.Println(" Min Delta : ", minDelt, minI+1)
	n := len(x) + 1
	A := 0.0
	B := 0.0
	C := 0.0
	D := 0.0
	for i := 0; i < n-1; i++ {
		A += math.Log(x[i]) * math.Log(x[i])
		B += math.Log(x[i])
		C += y[i] * math.Log(x[i])
		D += y[i]
	}
	a := (C - D*B/float64(n)) / (A - B*B/float64(n))
	b := (D - a*B) / float64(n)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	var z [9]float64
	for i, xx := range x {
		// z6
		z[i] = xx / (a*xx + b)
	}
	fmt.Println("z = ", z)
	del := 0.0
	for i := 0; i < n-1; i++ {
		del += (z[i] - y[i]) * (z[i] - y[i]) / float64(n)
	}
	del = math.Sqrt(del)
	fmt.Println(" del = ", del)
}

// итог z6
