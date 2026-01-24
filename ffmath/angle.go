// SPDX-FileCopyrightText: 2026 Kalle Fagerberg
// SPDX-FileCopyrightText: 2014-present Godot Engine contributors (see AUTHORS.md: https://github.com/godotengine/godot/blob/4.5.1-stable/AUTHORS.md)
// SPDX-FileCopyrightText: 2007-2014 Juan Linietsky, Ariel Manzur
//
// SPDX-License-Identifier: MIT

package ffmath

import (
	"math"

	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

// Linearly interpolates between two angles by the factor defined in "weight".
//
// Weight should be between 0.0 and 1.0 (inclusive).
// However, values outside this range are allowed and can be used to perform
// extrapolation.
// If this is not desired then you can use [Clamp01] to limit the weight.
func LerpAngle(from, to firefly.Angle, weight float32) firefly.Angle {
	return firefly.Radians(from.Radians() + AngleDifference(from, to).Radians()*weight)
}

// Angle difference to go from "a" to "to".
//
// Result will be in the range of [-[math.Pi], +[math.Pi]].
// When "a" and "to" are opposite,
// returns -[math.Pi] if "a" is smaller than "to", or [math.Pi] otherwise.
//
// Input angles do not need to be normalized.
//
// Based on the Godot [angle_difference] (licensed under MIT)
//
// [angle_difference]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L482-L489
func AngleDifference(from, to firefly.Angle) firefly.Angle {
	// tinymath has "RemEuclid" https://github.com/orsinium-labs/tinymath/blob/v1.1.0/tinymath.go#L324-L332
	// but it has some math bugs, so we have to resort to the big math functions.
	diff := Mod(to.Radians()-from.Radians(), 2*math.Pi)
	return firefly.Radians(Mod(2*diff, 2*math.Pi) - diff)
}

// Rotates "a" toward "to" by the "delta" amount.
//
// Will not go past "to", but interpolated correctly when the angles
// wrap around [Radians](2*[math.Pi]) or [Degrees](360).
//
// If "delta" is negative, this function will rotate away from "to",
// towards the opposite angle, and will not go past the opposite angle.
//
// Based on the Godot [rotate_towards] (licensed under MIT)
//
// [rotate_towards]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L598-L609
func RotateTowards(from, to, delta firefly.Angle) firefly.Angle {
	diff := AngleDifference(from, to).Radians()
	absDiff := tinymath.Abs(diff)
	return firefly.Radians(
		from.Radians() + Clamp(delta.Radians(), absDiff-math.Pi, absDiff)*tinymath.Sign(diff),
	)
}
