package main

import (
	"fmt"
	"math"
)

func calculateMean(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func calculateStandardDeviation(numbers []float64) float64 {
	mean := calculateMean(numbers)
	var variance float64
	for _, num := range numbers {
		variance += math.Pow(num-mean, 2)
	}
	variance /= float64(len(numbers))
	return math.Sqrt(variance)
}

func main() {
	numbers := []float64{10, 12, 23, 23, 16, 23, 21, 16}
	standardDeviation := calculateStandardDeviation(numbers)
	fmt.Printf("Standard Deviation: %.2f\n", standardDeviation)
}