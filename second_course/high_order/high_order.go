package main

import (
	"fmt"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * (t * t)) + (v * t) + s
	}
	return fn
}

func main() {
	fmt.Println("Enter the value of acceleration:")
	var acceleration float64
	fmt.Scan(&acceleration)

	fmt.Println("Enter the value of initial velocity:")
	var initialVelocity float64
	fmt.Scan(&initialVelocity)

	fmt.Println("Enter the value of initial displacement:")
	var initialDisplacement float64
	fmt.Scan(&initialDisplacement)

	calcDisplacement := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println("Enter the value of time for displacement calculation:")
	var time float64
	fmt.Scan(&time)

	fmt.Printf("Displacement equals to %g", calcDisplacement(time))

}
