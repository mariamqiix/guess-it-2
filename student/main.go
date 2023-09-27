package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	x := 0
	var array []int
	for scanner.Scan() {
		currentY, _ := strconv.Atoi(scanner.Text())
		x++
		array = append(array, currentY)
		a, b := CorrelationCoefficent(array,x)
		y2 := (a * float64(x)) + b
		y := y2 - float64(currentY)
		fmt.Printf("%d %d\n", int(float64(currentY)-y), int(y2))
	}
}

func ava(i []int, no int) float64 {
	z := 0.0
	if no < 10 {
		for x := 0; x < len(i); x++ {
			z += float64(i[x])
		}
	} else {
		for x := no - 10; x < no; x++ {
			z += float64(i[x])
		}
	}
	return (z / float64(10))
}

func Variance(i []int, no int) float64 {
	DeltaX := ava(i,no)
	TheX := 0.0
	Variance := 0.0
	if no <= 10 {
		for x := 0; x < len(i); x++ {
			z := (float64(i[x]) - float64(DeltaX))
			TheX += (z * z)
		}
		Variance = TheX / (float64(len(i)))
	} else {
		for x := no - 10; x < no; x++ {
			z := (float64(i[x]) - float64(DeltaX))
			TheX += (z * z)
		}
		Variance = TheX / (float64(10))
	}

	return Variance
}

func StandardDeviation(i []int,no int) float64 {
	x := Variance(i,no)
	return (math.Sqrt(x))
}

func CorrelationCoefficent(i []int, no int) (float64, float64) {
	theX := 0.0
	theY := 0.0
	theXY := 0.0
	theXX := 0.0
	theYY := 0.0
	SXY := 0.0
	SXX := 0.0
	if no <= 10 {
		for x := 0; x < len(i); x++ {
			theXY += float64((x + 1) * i[x])
			theX += float64(x + 1)
			theY += float64(i[x])
			theXX += float64((x + 1) * (x + 1))
			theYY += float64(i[x] * i[x])
		}
		r := math.Round((((theY*theXX)-(theX*theXY))/((float64(len(i))*theXX)-(theX*theX)))*1000000) / 1000000
		SXY = theXY - ((theY * theX) / float64(len(i)))
		SXX = theXX - ((theX * theX) / float64(len(i)))
		a := SXY / SXX
		return a, r
	} else {
		for x := no - 10; x < no; x++ {
			theXY += float64((x + 1) * i[x])
			theX += float64(x + 1)
			theY += float64(i[x])
			theXX += float64((x + 1) * (x + 1))
			theYY += float64(i[x] * i[x])
		}
		r := math.Round((((theY*theXX)-(theX*theXY))/((float64(10)*theXX)-(theX*theX)))*1000000) / 1000000
		SXY = theXY - ((theY * theX) / float64(10))
		SXX = theXX - ((theX * theX) / float64(10))
		a := SXY / SXX
		return a, r
	}
return 0.0,0.0
}
