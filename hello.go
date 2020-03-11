package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return lim
	}
	return lim
}

// Sqrt exported comment
func Sqrt(x float64) float64 {
	z := 1.0
	lastZ := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("trying:", i, z)
		if math.Abs(lastZ-z) < 0.0001 {
			break
		} else {
			lastZ = z
		}
	}
	return z
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println(Sqrt(5))

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

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
