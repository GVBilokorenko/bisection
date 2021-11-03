package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func polynom(x float64, coeffs []float64) float64 {
	var res = 0.0
	for i,coef := range coeffs {
		 var n = len(coeffs) - i - 1
		 var add = coef * math.Pow(x, float64(n))
		 res += add
	}
	return res
}

func Bisection(lower float64, upper float64, acc float64, coeffs []float64) (error, float64) {
	if polynom(lower, coeffs) * polynom(upper, coeffs) >= 0 {
		return errors.New("You have not assumed right lower and upper bound\n"), 0.0
	}
	var c float64 = lower
	for math.Abs(upper - lower) >= acc {
		c = (lower + upper) / 2
		if polynom(c, coeffs) == 0 {break}
		if (polynom(c, coeffs) * polynom(upper, coeffs) < 0) {
			lower = c
		} else {
			upper = c
		}
	}
	return nil, c
}

var (
	upper = flag.Float64("u", 0, "upper bound")
	lower = flag.Float64("l", 0, "lower bound")
	acc = 	flag.Float64("a", 0, "upper bound")
	coeffs = flag.String("c", "", "coefficients")
)

func main() {
	flag.Parse()

	if *upper == *lower {
		fmt.Println(os.Stderr, "upper and lower bounds are equal")
		return
	}

	if *acc == 0 {
		fmt.Println(os.Stderr, "accuracy should be different from 0")
		return
	}

	if len(*coeffs) == 0 {
		fmt.Println(os.Stderr, "coefficients are missing")
		return
	}

	cfsFloat64 := strings.Fields(*coeffs);
	cfs := []float64{}

	for _, arg := range cfsFloat64 {
        if n, err := strconv.ParseFloat(arg, 64); err == nil {
            cfs = append(cfs, n)
        }
    }

	var _, res = Bisection(*lower, *upper, *acc, cfs)
	fmt.Fprintln(os.Stdout, res)
	
}
