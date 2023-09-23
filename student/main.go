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
		if currentY < 50 {
			currentY = 100
		} else if currentY > 350 {
			currentY = 200
		}
		array = append(array, currentY)
		y := math.Round(ava(array))
		standard := math.Round(StandardDeviation(array))
		a,b := CorrelationCoefficent(array)
		y2 := (a*x)+b
		fmt.Printf("%d %d\n", (int((y - standard))+y2/2),(int((y + standard))+y2/2))
	}
}

func ava(i []int) float64 {
	z := 0.0
	for x := 0; x < len(i); x++ {
		z += float64(i[x])
	}
	return (z / float64(len(i)))
}

func Variance(i []int) float64 {
	DeltaX := ava(i)
	TheX := 0.0
	Variance := 0.0
	for x := 0; x < len(i); x++ {
		z := (float64(i[x]) - float64(DeltaX))
		TheX += (z * z)
	}
	Variance = TheX / (float64(len(i)))
	return Variance
}

func StandardDeviation(i []int) float64 {
	x := Variance(i)
	return (math.Sqrt(x))
}


func CorrelationCoefficent(i []int) (float64,float64) {
	theX := 0.0
	theY := 0.0
	theXY := 0.0
	theXX := 0.0
	theYY := 0.0
	SXY := 0.0
	SXX := 0.0
	for x := 0; x < len(i); x++ {
		theXY += float64((x + 1) * i[x])
		theX += float64(x + 1)
		theY += float64(i[x])
		theXX += float64((x + 1) * (x + 1))
		theYY += float64(i[x] * i[x])
	}
	r := 	d := math.Round((((theY*theXX)-(theX*theXY))/((float64(len(i))*theXX)-(theX*theX)))*1000000) / 1000000
	SXY = theXY - ((theY * theX) / float64(len(i)))
	SXX = theXX - ((theX * theX) / float64(len(i)))
	a := SXY / SXX
	return a,r
}