// SPDX-FileCopyrightText: 2026 Kalle Fagerberg
// SPDX-FileCopyrightText: 2014-present Godot Engine contributors (see AUTHORS.md: https://github.com/godotengine/godot/blob/4.5.1-stable/AUTHORS.md)
// SPDX-FileCopyrightText: 2007-2014 Juan Linietsky, Ariel Manzur
//
// SPDX-License-Identifier: MIT

// Package ffmath provides some utility math functions commonly used
// in game programming.
package ffmath

import (
	"cmp"
	"math"

	"github.com/orsinium-labs/tinymath"
)

const (
	// Multiply a number by this factor to convert it from a radian to a degree.
	RadToDeg = 360 / tinymath.Tau
	// Multiply a number by this factor to convert it from a degree to a radian.
	DegToRad = tinymath.Tau / 360
	// Smallest rounding error used in functions like [EqualApprox] and [IsZeroApprox]
	Epsilon = 0.00001
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

// Returns value clamped between minimum and maximum.
//
//   - If value is less than minimum, then you get minimum
//   - If value is more than maximum, then you get maximum
func Clamp[T cmp.Ordered](val, minimum, maximum T) T {
	switch {
	case val < minimum:
		return minimum
	case val > maximum:
		return maximum
	default:
		return val
	}
}

// Returns value clamped between 0 and 1.
//
//   - If value is less than 0, then you get 0
//   - If value is more than 1, then you get 1
func Clamp01[T Number](val T) T {
	switch {
	case val < 0:
		return 0
	case val > 1:
		return 1
	default:
		return val
	}
}

// Moves "start" towards "end" by "delta" amount.
//
// Returned value will not go past "end".
//
// Use a negative "delta" value to move away from "end".
//
// Based on the Godot [move_toward] (licensed under MIT)
//
// [move_toward]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L591-L596
func MoveTowards[T Number](start, end, delta T) T {
	if Abs(end-start) <= delta {
		return end
	} else {
		return start + Sign(end-start)*delta
	}
}

// Linear interpolation between two values by the factor defined in "weight".
//
// Weight should be between 0.0 and 1.0 (inclusive).
// However, values outside this range are allowed and can be used to perform
// extrapolation.
// If this is not desired then you can use [Clamp01] to limit the weight.
//
// See also [InverseLerp] that performs the reverse of this operation.
//
// Based on the Godot [lerp] (licensed under MIT)
//
// [lerp]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L336-L341
func Lerp[T Number](from, to, weight T) T {
	return from + (to-from)*weight
}

// Performs a reverse [Lerp], returning the weight factor of the value in the range.
//
//   - Return is between [0, 1] if the value is between [from, to]
//   - Return <0 if the value is below 'from'
//   - Return >1 if the value is above 'to'
//
// Based on the Godot [inverse_lerp] (licensed under MIT)
//
// [inverse_lerp]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L498-L503
func InverseLerp[T Number](from, to, value T) T {
	return (value - from) / (to - from)
}

// Wraps float32 in the half-open range [min, max) by wrapping around instead of clamping.
//
// Using Wrap with min=0 is equivalent to using [Mod], so prefer using that
// as Wrap is even more computationally expensive.
//
// Based on the Godot [wrap] (licensed under MIT)
//
// [wrap]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L632-L653
func Wrap[T Number](value, min, max T) T {
	delta := max - min
	if IsZeroApprox(delta) {
		return min
	}
	result := value - (delta * Floor((value-min)/delta))
	if EqualApprox(result, max) {
		return max
	}
	return result
}

// Check if two numbers are approximately equal to each other.
//
// The comparison done here is to see if the difference between the numbers
// are less than [Epsilon].
//
// Infinite values with the same sign (+/-) are considered equal.
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
//
// Based on the Godot [is_equal_approx] (licensed under MIT)
//
// [is_equal_approx]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L512-L552
func EqualApprox[T Number](a, b T) bool {
	switch x := any(a).(type) {
	case float32:
		if a == b {
			return true
		}
		tolerance := Epsilon * tinymath.Abs(x)
		if tolerance < Epsilon {
			tolerance = Epsilon
		}
		return tinymath.Abs(x-float32(b)) < tolerance
	case float64:
		if a == b {
			return true
		}
		tolerance := Epsilon * math.Abs(x)
		if tolerance < Epsilon {
			tolerance = Epsilon
		}
		return math.Abs(x-float64(b)) < tolerance
	default:
		// all other types are integers
		return a == b
	}
}

// Check if a numbers is approximately equal to zero.
//
// The comparison done here is to see if the difference between the numbers
// are less than [Epsilon].
//
// This function is faster than running [EqualApprox](a, 0)
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
//
// Based on the Godot [is_zero_approx] (licensed under MIT)
//
// [is_zero_approx]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L554-L559
func IsZeroApprox[T Number](a T) bool {
	switch x := any(a).(type) {
	case float32:
		return tinymath.Abs(x) < Epsilon
	case float64:
		return math.Abs(x) < Epsilon
	default:
		// all other types are integers
		return a == 0
	}
}

// Utility function to calculate "lhs % rhs" on float32.
//
// The [tinymath] library also has [tinymath.RemEuclid], which is supposed to
// fulfill the same operation. But it can give somewhat unstable results
// in some instances.
//
// This function relies on [math.Mod] by converting the float32 to float64
// and back, which is more computationally intensive on 32-bit machines like the
// Firefly Zero. Only use this as a last resort.
func Mod(lhs, rhs float32) float32 {
	// tinymath has "RemEuclid" https://github.com/orsinium-labs/tinymath/blob/v1.1.0/tinymath.go#L324-L332
	// but it has some math bugs, so we have to resort to the big math functions.
	return float32(math.Mod(float64(lhs), float64(rhs)))
}

// Generic function for getting the floored value of a number.
//
// Under the hood the function uses different code paths for different types:
//
//   - float32: [tinymath.Floor]
//   - float64: [math.Floor]
//   - integers: returns the value as-is
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
func Floor[T Number](a T) T {
	switch x := any(a).(type) {
	case float32:
		return T(tinymath.Floor(x))
	case float64:
		return T(math.Floor(x))
	default:
		// all other types are integers
		return a
	}
}

// Generic function for getting the ceiling value of a number.
//
// Under the hood the function uses different code paths for different types:
//
//   - float32: [tinymath.Ceil]
//   - float64: [math.Ceil]
//   - integers: returns the value as-is
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
func Ceil[T Number](a T) T {
	switch x := any(a).(type) {
	case float32:
		return T(tinymath.Ceil(x))
	case float64:
		return T(math.Ceil(x))
	default:
		// all other types are integers
		return a
	}
}

// Generic function for getting the rounded value of a number.
//
// Under the hood the function uses different code paths for different types:
//
//   - float32: [tinymath.Round]
//   - float64: [math.Round]
//   - integers: returns the value as-is
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
func Round[T Number](a T) T {
	switch x := any(a).(type) {
	case float32:
		return T(tinymath.Round(x))
	case float64:
		return T(math.Round(x))
	default:
		// all other types are integers
		return a
	}
}

// Generic function for getting the absolute value of a number.
//
// Under the hood the function uses different code paths for different types:
//
//   - float32: [tinymath.Abs]
//   - float64: [math.Abs]
//   - unsigned integer: returns the value as-is
//   - signed integer: -a if a<0, a otherwise
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
func Abs[T Number](a T) T {
	switch x := any(a).(type) {
	case float32:
		return T(tinymath.Abs(x))
	case float64:
		return T(math.Abs(x))
	case uint, uintptr, uint8, uint16, uint32, uint64:
		// unsigned, can't be negative
		return a
	default:
		// all other types are signed integers
		if a < 0 {
			return -a
		}
		return a
	}
}

// Generic function for getting the sign, i.e +1 if positive, -1 if negative, 0 if zero.
//
// Under the hood the function uses different code paths for different types:
//
//   - float32: [tinymath.Sign]
//   - float64: [math.Copysign]
//   - unsigned integers: 1 if a>0, 0 otherwise
//   - signed integers: 1 if a>0, -1 if a<0, 0 otherwise
//
// This function is generic just as a utility so it can be used in conjunction
// with other generic functions from this package.
func Sign[T Number](a T) T {
	switch x := any(a).(type) {
	case float32:
		return T(tinymath.Sign(x))
	case float64:
		return T(math.Copysign(1.0, x))
	case uint, uintptr, uint8, uint16, uint32, uint64:
		// unsigned, can't be negative
		if a == 0 {
			return 0
		}
		return 1
	default:
		// all other types are signed integers
		if a < 0 {
			return any(-1).(T)
		}
		if a > 0 {
			return any(1).(T)
		}
		return a
	}
}

// Returns true if the float is neither NaN nor infinity.
func IsFinite(f float32) bool {
	return !tinymath.IsNaN(f) && f > tinymath.NegInf && f < tinymath.Inf
}
