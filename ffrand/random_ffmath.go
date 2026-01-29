// SPDX-FileCopyrightText: 2026 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package ffrand

import "github.com/applejag/firefly-go-math/ffmath"

// Pseudo-random unit vector, where the vector's radious will be 1.
func VecUnit() ffmath.Vec {
	return ffmath.VAngle(Angle())
}

// Pseudo-random Vec in the half-open interval [min, max)
func VecRange(min, max ffmath.Vec) ffmath.Vec {
	return ffmath.V(Float32Range(min.X, max.X), Float32Range(min.Y, max.Y))
}
