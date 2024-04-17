package main

import (
	"fmt"
	"math"
)

func f(x1 float64, x2 float64) float64 {
	return x1 + 2*x2 + 4*math.Pow(1+math.Pow(x1, 2)+math.Pow(x2, 2), 0.5)
}

func derX1(x1 float64, x2 float64) float64 {
	return ((4 * x1) / (math.Pow(math.Pow(x1, 2)+math.Pow(x2, 2)+1, 0.5))) + 1
}

func derX2(x1 float64, x2 float64) float64 {
	return ((4 * x2) / (math.Pow(math.Pow(x1, 2)+math.Pow(x2, 2)+1, 0.5))) + 2
}
func derX1x1(x1 float64, x2 float64) float64 {
	return (4 * (math.Pow(x2, 2) + 1)) / math.Pow(math.Pow(x1, 2)+math.Pow(x2, 2)+1, 1.5)

}

func derX1x2(x1 float64, x2 float64) float64 {
	return (4 * x1 * x2) / math.Pow(math.Pow(x1, 2)+math.Pow(x2, 2)+1, 1.5)
}

func derX2x2(x1 float64, x2 float64) float64 {
	return (4 * (math.Pow(x1, 2) + 1)) / math.Pow(math.Pow(x1, 2)+math.Pow(x2, 2)+1, 1.5)
}

func main() {
	eps := float64(0.001)
	k := 0
	x1k := 0.0
	x2k := 0.0
	max := float64(0.0)

	phi1 := -(derX1(x1k, x2k) * derX1(x1k, x2k)) - (derX2(x1k, x2k) * derX2(x1k, x2k))
	phi2 := derX1x1(x1k, x2k)*derX1(x1k, x2k)*derX1(x1k, x2k) +
		2*derX1x2(x1k, x2k)*derX1(x1k, x2k)*derX2(x1k, x2k) +
		derX2x2(x1k, x2k) + derX2(x1k, x2k)*derX2(x1k, x2k)
	t := -(phi1 / phi2)
	tmp1 := x1k
	tmp2 := x2k
	x1k = tmp1 - t*derX1(x1k, x2k)
	x2k = tmp2 - t*derX2(x1k, x2k)
	k++

	if math.Abs(derX1(x1k, x2k)) > math.Abs(derX2(x1k, x2k)) {
		max = math.Abs(derX1(x1k, x2k))
	} else {
		max = math.Abs(derX2(x1k, x2k))
	}

	for max >= eps {
		phi1 := -(derX1(x1k, x2k) * derX1(x1k, x2k)) - (derX2(x1k, x2k) * derX2(x1k, x2k))
		phi2 := derX1x1(x1k, x2k)*derX1(x1k, x2k)*derX1(x1k, x2k) +
			2*derX1x2(x1k, x2k)*derX1(x1k, x2k)*derX2(x1k, x2k) +
			derX2x2(x1k, x2k)*derX2(x1k, x2k)*derX2(x1k, x2k)
		t := -(phi1 / phi2)
		tmp1 := x1k
		tmp2 := x2k
		x1k = tmp1 - t*derX1(tmp1, tmp2)
		//fmt.Println("", x1k)
		x2k = tmp2 - t*derX2(tmp1, tmp2)
		//fmt.Println(x2k)
		k++

		if math.Abs(derX1(x1k, x2k)) > math.Abs(derX2(x1k, x2k)) {
			max = math.Abs(derX1(x1k, x2k))
		} else {
			max = math.Abs(derX2(x1k, x2k))
		}
	}
	fmt.Println("minimum:")
	fmt.Println(x1k, " ", x2k)
	fmt.Println("analitical minimum:")
	fmt.Println(-1/math.Sqrt(11), " ", -2/math.Sqrt(11))
	//fmt.Println("difference:")
	//fmt.Println(math.Abs(x1k-(-1/m
	//ath.Sqrt(11))), " ", math.Abs(x2k-(-2/math.Sqrt(11))))
	fmt.Println("count iterations: ", k)
}
