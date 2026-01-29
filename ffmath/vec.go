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

// Utility type for dealing with float-based positions.
type Vec struct {
	X float32
	Y float32
}

// Shortcut for creating a [Vec].
func V(x, y float32) Vec {
	return Vec{X: x, Y: y}
}

// Converts a [firefly.Point] to a [Vec]
func VPoint(point firefly.Point) Vec {
	return Vec{X: float32(point.X), Y: float32(point.Y)}
}

// Converts a [firefly.Angle] to a [Vec]
func VAngle(angle firefly.Angle) Vec {
	return Vec{
		X: tinymath.Cos(angle.Radians()),
		Y: -tinymath.Sin(angle.Radians()),
	}
}

// Convert a [Vec] to a [Point].
//
// The X and Y floats are truncated, meaning the floored value of positive numbers
// and the ceiling value of negative values.
func (v Vec) Point() firefly.Point {
	return firefly.Point{X: int(v.X), Y: int(v.Y)}
}

// Return a vector with absolute X and Y components.
func (v Vec) Abs() Vec {
	return Vec{X: tinymath.Abs(v.X), Y: tinymath.Abs(v.Y)}
}

// Adds a position.
func (v Vec) Add(rhs Vec) Vec {
	return Vec{X: v.X + rhs.X, Y: v.Y + rhs.Y}
}

// Subtract a position.
func (v Vec) Sub(rhs Vec) Vec {
	return Vec{X: v.X - rhs.X, Y: v.Y - rhs.Y}
}

// Get a position with -X and -Y
func (v Vec) Negate() Vec {
	return Vec{X: -v.X, Y: -v.Y}
}

// Get a position with both X and Y to their minimum in the two given [Vec]s.
func (v Vec) ComponentMin(r Vec) Vec {
	if r.X < v.X {
		v.X = r.X
	}
	if r.Y < v.Y {
		v.Y = r.Y
	}
	return v
}

// Get a position with both X and Y to their maximum in the two given [Vec]s.
func (v Vec) ComponentMax(r Vec) Vec {
	if r.X > v.X {
		v.X = r.X
	}
	if r.Y > v.Y {
		v.Y = r.Y
	}
	return v
}

func (v Vec) Clamp(min, max Vec) Vec {
	return Vec{Clamp(v.X, min.X, max.X), Clamp(v.Y, min.Y, max.Y)}
}

// Get a position with both X and Y rounded to the nearest integer.
func (v Vec) Round() Vec {
	return Vec{X: tinymath.Round(v.X), Y: tinymath.Round(v.Y)}
}

// Get a position with both X and Y ceiled to the nearest integer.
func (v Vec) Ceil() Vec {
	return Vec{X: tinymath.Ceil(v.X), Y: tinymath.Ceil(v.Y)}
}

// Get a position with both X and Y floored to the nearest integer.
func (v Vec) Floor() Vec {
	return Vec{X: tinymath.Floor(v.X), Y: tinymath.Floor(v.Y)}
}

// Check if the position is within the screen boundaries.
func (v Vec) InBounds() bool {
	return v.X >= 0 && v.Y >= 0 && v.X < firefly.Width && v.Y < firefly.Height
}

// Get a position where both the X and Y value are individually multiplied by the scalar factor.
func (v Vec) Scale(factor float32) Vec {
	return Vec{X: v.X * factor, Y: v.Y * factor}
}

// Radius returns the vector length (aka magnitude).
//
// Uses [tinymath] for faster but less accurate calculation, with an average deviation of ~5%.
func (v Vec) Radius() float32 {
	return tinymath.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Get the squared vector length (aka squared magnitude), which is simpler to calculate.
func (v Vec) RadiusSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// The angle of the polar coordinate of the vector.
//
//   - [V](1, 0).Azimuth() == [firefly.Degrees](0)
//   - [V](0, 1).Azimuth() == [firefly.Degrees](90)
//   - [V](-1, 0).Azimuth() == [firefly.Degrees](180)
//   - [V](0, -1).Azimuth() == [firefly.Degrees](270)
//
// Uses [tinymath] for faster but less accurate calculation, with error of `0.1620` degrees.
func (v Vec) Azimuth() firefly.Angle {
	r := math.Pi / 2. * tinymath.Atan2Norm(v.Y, v.X)
	return firefly.Radians(r)
}

// Get a position that has moved towards "to" by the "delta" amount, but will not go past "to".
//
// Use negative "delta" value to move away.
func (v Vec) MoveTowards(to Vec, delta float32) Vec {
	vd := to.Sub(v)
	dist := vd.Radius()
	if dist <= delta || dist < Epsilon {
		return to
	}
	return v.Add(vd.Scale(delta / dist))
}

// Get the distance to another position.
//
// Uses [tinymath] for faster but less accurate calculation, with an average deviation of ~5%.
func (v Vec) DistanceTo(to Vec) float32 {
	return v.Sub(to).Radius()
}

// Get the squared distance to another position, which is simpler to calculate.
func (v Vec) DistanceToSquared(to Vec) float32 {
	return v.Sub(to).RadiusSquared()
}

// Get a normalized vector.
//
// A normalized vector's [Vec.Radius] equals 1.
//
// Uses [tinymath] for faster but less accurate calculation, with an average deviation of ~5%.
func (v Vec) Normalize() Vec {
	squaredRadius := v.RadiusSquared()
	if squaredRadius == 0 {
		return Vec{}
	}
	radius := tinymath.Sqrt(squaredRadius)
	return Vec{
		X: v.X / radius,
		Y: v.Y / radius,
	}
}

// Check if the radius approximately equal to 1.
func (v Vec) IsNormalized() bool {
	return EqualApprox(v.RadiusSquared(), 1)
}

// Get the dot product of two vectors
func (v Vec) Dot(other Vec) float32 {
	return v.X*other.X + v.Y*other.Y
}

// Get the cross product of two vectors
func (v Vec) Cross(other Vec) float32 {
	return v.X*other.Y + v.Y*other.X
}

// True if the other vector has exactly the same float values.
//
// This is done by float equality, which is very sensitive due to
// floating point number precision errors.
func (v Vec) Equal(other Vec) bool {
	return v.X == other.X && v.Y == other.Y
}

// True if the other vector has approximately the same float values.
//
// The comparison done here is to see if the difference between the positions
// are less than [Epsilon].
//
// Infinite values with the same sign (+/-) are considered equal.
func (v Vec) EqualApprox(other Vec) bool {
	return EqualApprox(v.X, other.X) && EqualApprox(v.Y, other.Y)
}

// True if the vector is approximately equal to {0,0}.
//
// This function is faster than [Vec.EqualApprox]([V](0, 0)).
func (v Vec) IsZeroApprox() bool {
	return IsZeroApprox(v.X) && IsZeroApprox(v.Y)
}

// True if the position's X and Y are real numbers (not infinity and not NaN).
func (v Vec) IsFinite() bool {
	return IsFinite(v.X) && IsFinite(v.Y)
}
