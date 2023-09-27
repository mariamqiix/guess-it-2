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
		y2 := (a * float64(x+1)) + b
		fmt.Printf("%d %d\n", int(float64(y2)-(5.0)), int(y2+7.0))
	}
}

func CorrelationCoefficent(i []int, no int) (float64, float64) {
	theX := 0.0
	theY := 0.0
	theXY := 0.0
	theXX := 0.0
	theYY := 0.0
	SXY := 0.0
	SXX := 0.0
	// if no <= 5 {
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

}
