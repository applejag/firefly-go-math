// SPDX-FileCopyrightText: 2026 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package ffrand

import (
	"github.com/applejag/firefly-go-math/ffmath"
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

// Pseudo-random size.
//
// The returned point can be negative.
func Point() firefly.Point {
	return firefly.P(Int(), Int())
}

// Pseudo-random point in the half-open interval [0, n)
func Pointn(n firefly.Point) firefly.Point {
	return firefly.P(Intn(n.X), Intn(n.Y))
}

// Pseudo-random point in the half-open interval [min, max)
func PointRange(min, max firefly.Point) firefly.Point {
	return firefly.P(IntRange(min.X, max.X), IntRange(min.Y, max.Y))
}

// Pseudo-random size.
func Size() firefly.Size {
	return firefly.S(Int(), Int())
}

// Pseudo-random size in the half-open interval [0, n)
func Sizen(n firefly.Size) firefly.Size {
	return firefly.S(Intn(n.W), Intn(n.H))
}

// Pseudo-random size in the half-open interval [min, max)
func SizeRange(min, max firefly.Size) firefly.Size {
	return firefly.S(IntRange(min.W, max.W), IntRange(min.H, max.H))
}

// Pseudo-random angle in the half-open interval [0, τ)
//
// In other words, the range is:
//
//   - [0, 360°)
//   - [0, 2π)
func Angle() firefly.Angle {
	return firefly.Radians(Float32() * tinymath.Tau)
}

// Pseudo-random angle in the half-open interval [0, n)
func Anglen(n firefly.Angle) firefly.Angle {
	return firefly.Radians(Float32() * n.Normalize().Radians())
}

// Pseudo-random angle in the half-open interval [min, max)
func AngleRange(min, max firefly.Angle) firefly.Angle {
	return min.Add(firefly.Radians(Float32() * ffmath.AngleDifference(min, max).Radians()))
}
