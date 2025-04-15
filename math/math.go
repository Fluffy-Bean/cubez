// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

/*

The math module of this project defines the floating point precision to
be used and the mathematical types to be used.

From there it also defines mathematial operations on vectors, matrixes
and quaternions that operate by reference.

All matrixes will be created in column-major ordering.

*/

package math

import (
	"math"
)

// Espilon is used to test equality of the floats and represents how
// close two floats can be and still test positive for equality.
const Epsilon float64 = 1e-7

const (
	MinNormal = 1.1754943508222875e-38 // 1 / 2**(127 - 1)
	MinValue  = math.SmallestNonzeroFloat64
	MaxValue  = math.MaxFloat64
)

var (
	InfPos = math.Inf(1)
	InfNeg = math.Inf(-1)
	NaN    = math.NaN()
)

// Matrix3 is a 3x3 matrix of floats in column-major order.
type Matrix3 [9]float64

// Matrix3x4 is a 3x4 matrix of floats in column-major order
// that can be used to hold a rotation and translation in 3D space
// where the 4th row would have been [0 0 0 1]
type Matrix3x4 [12]float64

// Matrix4 is a 4x4 matrix of floats in column-major order.
type Matrix4 [16]float64

// FloatsEqual tests the two float64 numbers for equality, which really means
// that it tests whether or not the difference between the two is
// smaller than Epsilon.
func FloatsEqual(a, b float64) bool {
	// handle cases like inf
	if a == b {
		return true
	}

	diff := math.Abs(a - b)

	// if a or b is 0 or really close to it
	if a*b == 0 || diff < MinNormal {
		return diff < Epsilon*Epsilon
	}

	return diff/math.Abs(a)+math.Abs(b) < Epsilon
}

// DegToRad converts degrees to radians
func DegToRad(angle float64) float64 {
	return angle * math.Pi / 180.0
}

// RadToDeg converts radians to degrees
func RadToDeg(angle float64) float64 {
	return angle * 180.0 / math.Pi
}
