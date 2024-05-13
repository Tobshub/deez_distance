package main

import (
	"flag"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

const target = "deez"

func main() {
	var src string
	flag.StringVar(&src, "w", "nuts", "input string")
	flag.Parse()

	m, n := len(target), len(src)
	d := mat.NewDense(m+1, n+1, nil)

	for i := 1; i <= m; i++ {
		d.Set(i, 0, float64(i))
	}

	for j := 1; j <= n; j++ {
		d.Set(0, j, float64(j))
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			substitutionCost := 0
			if target[i-1] != src[j-1] {
				substitutionCost = 1
			}

			d.Set(i, j, min(
				d.At(i-1, j)+1,
				d.At(i, j-1)+1,
				d.At(i-1, j-1)+float64(substitutionCost),
			))
		}
	}

	fmt.Printf("Distance from Deez: %.2f\n", d.At(m, n))
}
