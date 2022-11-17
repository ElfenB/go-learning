package main

import (
	"fmt"
  "math"
	"runtime"
)

func Sqrt(x float64, tries int) float64 {
	// start with z=1
	z := 1.0

	for i := 0; i < tries; i++ {
		// if formula not 0, substract result
		if v := (z * z - x) / (2*z); v != 0 {
			z -= v
			fmt.Println("z:", z);
		}
		// repeat
	}

	return z
}

func ownSqrtFunc() {
	fmt.Println(Sqrt(2, 10))
	fmt.Println(math.Sqrt(2))
}

func printPlatform() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func learnPointers() {
	i := 3

	p := &i

	fmt.Println(*p);

	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	k := &v
	k.X = 2
	v.X = 4
	fmt.Println(v);
}

func learnArrays() {
	// Array with fixed size
	var a [5]int
	a[0] = 3
	fmt.Println(a);

	// Slice with variable size
	b := []int{1,2}
	fmt.Println(b);

	// Change value of slice that point to an array in the background
	b[0] = 2
	fmt.Println(b);
}

func main() {
	// ownSqrtFunc()

	// printPlatform()

	// learnPointers()

	learnArrays()
}
