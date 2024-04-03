package main

import (
	"fmt"
	"math"
)

func run_through_method_1(a, b, c, d []float64) []float64 {
	alpha := []float64{-c[0] / b[0]}
	beta := []float64{d[0] / b[0]}
	n := len(d)
	x := make([]float64, n)

	for i := 1; i < n-1; i++ {
		alpha = append(alpha, -c[i]/(b[i]+a[1]*alpha[i-1]))
		beta = append(beta, (d[i]-a[i-1]*beta[i-1])/(b[i]+a[i-1]*alpha[i-1]))
	}

	beta = append(beta, (d[n-1]-a[n-2]*beta[n-2])/(b[n-1]+a[n-2]*alpha[n-2]))

	x[n-1] = (d[n-1] - a[n-2]*beta[n-2]) / (a[n-2]*alpha[n-2] + b[n-1])

	for i := n - 1; i > 0; i-- {
		x[i-1] = alpha[i-1]*x[i] + beta[i-1]
	}

	return x
}

func main() {
	p := 5.0
	q := -4.0
	n := 10
	xa := 0.0

	xb := 1.0
	h := (xb - xa) / float64(n)
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	for i := 0; i < n+1; i++ {
		x[i] = float64(i) / 10.0
	}
	//fmt.Println(float64(n+1)*h, h, n/10)
	y[0] = math.Exp(0)
	fmt.Println(y[0])
	y[n] = math.Exp(1)
	fmt.Println(y[n])

	//Мето
	a := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = 1 - p*h/2.0
	}
	b := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = math.Pow(h, 2)*q - 2
	}
	c := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		c[i] = 1 + p*h/2.0
	}

	d := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		d[i] = math.Pow(h, 2) * math.Exp(x[i])
	}
	d[0] -= 1 * a[0]
	d[n-2] -= math.E * c[0]

	solution := run_through_method_1(a, b, c, d)
	fmt.Println("i\tx\t\ty\t\ty'\t\tf-y")
	//solution:=run_through_method_1(N, free)
	for i := 1; i < n; i++ {
		y[i] = solution[i-1]
	}
	//fmt.Println(y)
	for i := 0; i < n+1; i++ {
		//fmt.Println(N[i])
		yy := math.Exp(x[i])
		fmt.Printf("%d\t%f\t%f\t%f\t%f\n", i, x[i], yy, y[i], math.Abs(yy-y[i]))
	}
	// y'' + py' +qy = 2e^x
	// y(x) = e^x,
}
