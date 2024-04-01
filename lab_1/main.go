package main

import (
	"fmt"
	"math"
)

// че та проверить в узлах интерполяции (щас высчитывается между ними ) - последний столбик, там д.б. дроби меньше
func run_through_method(A [30][30]float64, d []float64) []float64 {
	var alpha, beta, x []float64
	//прямой ход
	for i := 0; i < len(A); i++ {
		if i == 0 {
			alph := -1 * A[i][i+1] / A[i][i]
			bet := d[i] / A[i][i]
			alpha = append(alpha, alph)
			beta = append(beta, bet)
		} else if i == len(A)-1 {
			alpha = append(alpha, 0)
			bet := (d[i] - A[i][i-1]*beta[i-1]) / (A[i][i-1]*alpha[i-1] + A[i][i])
			beta = append(beta, bet)
		} else {
			alph := -1 * A[i][i+1] / (A[i][i-1]*alpha[i-1] + A[i][i])
			bet := (d[i] - A[i][i-1]*beta[i-1]) / (A[i][i-1]*alpha[i-1] + A[i][i])
			alpha = append(alpha, alph)
			beta = append(beta, bet)
		}

	}

	//обратный ход
	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			temp := beta[i]
			x = append(x, temp)
		} else {
			temp := alpha[i]*x[len(x)-1] + beta[i]
			x = append(x, temp)
		}
	}

	var right_x []float64

	for i := len(x) - 1; i >= 0; i-- {
		right_x = append(right_x, x[i])
	}
	right_x[0] = 0
	right_x[len(right_x)-1] = 0
	return right_x
}

func tab(alph, bet float64) {
	n := 32
	h := (bet - alph) / float64(n)
	var x, y, x1, y1 []float64
	for i := alph; i <= bet; i += h {
		x = append(x, i)
		temp := i / (math.Pow(i, 2) - 2*i - 3)
		y = append(y, temp)

	}

	var a, b, d, c []float64

	//подготовка к методу прогонки
	var left []float64
	var mat [30][30]float64
	for i := 1; i < 31; i++ {
		left = append(left, (3/math.Pow(h, 2))*(y[i+1]-2*y[i]+y[i-1]))
	}
	for i := 0; i < (n - 2); i++ {
		for j := 0; j < 30; j++ {
			if i == j {
				mat[i][j] = 4
			} else if (i-j == 1) || (j-i == 1) {
				mat[i][j] = 1
			}
		}
	}
	c = append(c, 0)
	c = append(c, run_through_method(mat, left)...)
	c = append(c, 0)

	for i := 0; i < n-1; i++ {
		a = append(a, y[i])
		b = append(b, (y[i+1]-y[i])/h-(h/3)*(c[i+1]+2*c[i]))
		d = append(d, (c[i+1]-c[i])/(3*h))
	}

	for i := 1; i < n; i++ {
		xx := alph + (float64(i)-0.5)*h
		x1 = append(x1, xx)
		temp := xx * math.Pow(math.Sin(xx), 2)
		y1 = append(y1, temp)
	}

	var f, ff []float64

	for i := 0; i < n-1; i++ {
		f = append(f, a[i]+b[i]*(x1[i]-x[i])+c[i]*math.Pow(x1[i]-x[i], 2)+d[i]*math.Pow(x1[i]-x[i], 3))
		ff = append(ff, math.Abs(f[i]-y[i]))
	}

	fmt.Println("i \t x\t\t f\t\t y\t\t |f-y|")
	for i := 0; i < n-1; i++ {
		fmt.Printf("%d:\t%f\t%f\t%f\t%f\n", i, x[i], f[i], y[i], ff[i])
	}

}

func main() {
	tab(-0.5, 2.5)
}
