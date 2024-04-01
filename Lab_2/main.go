package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x / (math.Pow(x, 2) - 2*x - 3)
}

func midRectangle(x []float64, h float64, n int) float64 {
	integ := float64(0)
	for i := 0; i < n; i++ {
		integ += f(x[i] + h*0.5)
	}
	return h * integ
}
func trap(x []float64, h float64, n int) float64 {

	mid := float64(0)
	for i := 1; i < n; i++ {
		mid += f(x[i])
	}
	return h * (mid + ((f(x[0]) + f(x[n-1])) / 2))

}

func simps(x []float64, h float64, n int) float64 {
	sec, third := float64(0), float64(0)
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			third += f(x[i])
		} else {
			sec += f(x[i])
		}

	}

	return h * (f(x[0]) + f(x[n-1]) + 4*sec + 2*third) / 3
}

func rich(f1, f2 float64, k int) float64 {
	return (f1 - f2) / (math.Pow(2.0, float64(k)) - 1.0)
}

func main() {
	a := -0.5
	b := 2.5

	eps := 0.001
	fmt.Printf("eps: %f\n", eps)
	rich1 := 1.1
	int1 := 0.0
	n1 := 1
	for math.Abs(rich1) > eps {
		n1 *= 2
		h1 := (b - a) / float64(n1)
		var x1 []float64 = make([]float64, n1)
		for i := 0; i < n1; i++ {
			x1[i] = a + float64(i)*h1
		}

		temp := int1
		int1 = midRectangle(x1, h1, n1)

		rich1 = rich(int1, temp, 2)
	}
	rich2 := 1.1
	n2 := 1
	int2 := 0.0
	for math.Abs(rich2) > eps {
		n2 *= 2
		h1 := (b - a) / float64(n2)
		var x1 []float64 = make([]float64, n2)
		for i := 0; i < n2; i++ {
			x1[i] = a + float64(i)*h1
		}

		temp := int2
		int2 = trap(x1, h1, n2)

		rich2 = rich(int2, temp, 2)
	}

	rich3 := 1.1
	n3 := 1
	int3 := 0.0
	for math.Abs(rich3) > eps {
		n3 *= 2
		h1 := (b - a) / float64(n3)
		var x1 []float64 = make([]float64, n3)
		for i := 0; i < n3; i++ {
			x1[i] = a + float64(i)*h1
		}

		temp := int3
		int3 = simps(x1, h1, n3)
		rich3 = rich(int3, temp, 2)
	}
	fmt.Println("\t\tMidRectangle \t\tTrapezoid \t\tSimpson")
	fmt.Printf("n\t\t%d\t\t\t%d\t\t\t%d\n", n1, n2, n3)
	fmt.Printf("I\t\t%f\t\t%f\t\t%f\n", int1, int2, int3)
	fmt.Printf("R\t\t%f\t\t%f\t\t%f\n", rich1, rich2, rich3)
	fmt.Printf("I+R\t\t%f\t\t%f\t\t%f\n", int1+rich1, int2+rich2, int3+rich3)

}
