// SPDX-FileCopyrightText: 2026 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package ffmath_test

import (
	"fmt"

	"github.com/applejag/firefly-go-math/ffmath"
)

func ExampleLerp() {
	fmt.Println("Lerp(0, 10, 0) =", ffmath.Lerp(0., 10., 0.))
	fmt.Println("Lerp(0, 10, 1) =", ffmath.Lerp(0., 10., 1.))
	fmt.Println("Lerp(0, 10, .5) =", ffmath.Lerp(0., 10., .5))
	fmt.Println("Lerp(0, 10, 2) =", ffmath.Lerp(0., 10., 2.))
	fmt.Println("Lerp(0, 10, -1) =", ffmath.Lerp(0., 10., -1.))

	// Output:
	// Lerp(0, 10, 0) = 0
	// Lerp(0, 10, 1) = 10
	// Lerp(0, 10, .5) = 5
	// Lerp(0, 10, 2) = 20
	// Lerp(0, 10, -1) = -10
}

func ExampleInverseLerp() {
	fmt.Println("InverseLerp(0, 10, 0) =", ffmath.InverseLerp(0., 10., 0.))
	fmt.Println("InverseLerp(0, 10, 10) =", ffmath.InverseLerp(0., 10., 10.))
	fmt.Println("InverseLerp(0, 10, 5) =", ffmath.InverseLerp(0., 10., 5.))
	fmt.Println("InverseLerp(0, 10, 20) =", ffmath.InverseLerp(0., 10., 20.))
	fmt.Println("InverseLerp(0, 10, -10) =", ffmath.InverseLerp(0., 10., -10.))

	// Output:
	// InverseLerp(0, 10, 0) = 0
	// InverseLerp(0, 10, 10) = 1
	// InverseLerp(0, 10, 5) = 0.5
	// InverseLerp(0, 10, 20) = 2
	// InverseLerp(0, 10, -10) = -1
}

func ExampleWrap() {
	var value float32 = 9
	for range 3 {
		nextValue := ffmath.Wrap(value+0.5, 5, 10)
		fmt.Printf("Wrap(%.1f+0.5, 5, 10) = %.1f\n", value, nextValue)
		value = nextValue
	}

	fmt.Println()
	for range 3 {
		nextValue := ffmath.Wrap(value-0.5, 5, 10)
		fmt.Printf("Wrap(%.1f-0.5, 5, 10) = %.1f\n", value, nextValue)
		value = nextValue
	}

	// Output:
	// Wrap(9.0+0.5, 5, 10) = 9.5
	// Wrap(9.5+0.5, 5, 10) = 5.0
	// Wrap(5.0+0.5, 5, 10) = 5.5
	//
	// Wrap(5.5-0.5, 5, 10) = 5.0
	// Wrap(5.0-0.5, 5, 10) = 9.5
	// Wrap(9.5-0.5, 5, 10) = 9.0
}
