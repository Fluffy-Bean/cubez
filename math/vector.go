// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package math

import (
	"math"
)

// Vector3 is a vector of three floats.
type Vector3 [3]float64

// Vector4 is a vector of four floats.
type Vector4 [4]float64

// Quat is the type of a quaternion (w,x,y,z).
type Quat Vector4

// Add adds a vector to another vector.
func (v *Vector3) Add(v2 *Vector3) {
	v[0] += v2[0]
	v[1] += v2[1]
	v[2] += v2[2]
}

// Add adds a vector, scaled by a float64, to another vector.
func (v *Vector3) AddScaled(v2 *Vector3, scale float64) {
	v[0] += v2[0] * scale
	v[1] += v2[1] * scale
	v[2] += v2[2] * scale
}

// Clear sets the vector to {0.0, 0.0, 0.0}.
func (v *Vector3) Clear() {
	v[0], v[1], v[2] = 0.0, 0.0, 0.0
}

// ComponentProduct performs a component-wise product with another vector.
func (v *Vector3) ComponentProduct(v2 *Vector3) {
	v[0] *= v2[0]
	v[1] *= v2[1]
	v[2] *= v2[2]
}

// Cross returns the cross product of this vector with another.
func (v *Vector3) Cross(v2 *Vector3) Vector3 {
	return Vector3{
		v[1]*v2[2] - v[2]*v2[1],
		v[2]*v2[0] - v[0]*v2[2],
		v[0]*v2[1] - v[1]*v2[0],
	}
}

// Dot returns the dot product of this vector with another.
func (v *Vector3) Dot(v2 *Vector3) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

// Magnitude returns the magnitude of the vector.
func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

// SquareMagnitude returns the magitude of the vector, squared.
func (v *Vector3) SquareMagnitude() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// MulWith multiplies a vector by a float64 number.
func (v *Vector3) MulWith(r float64) {
	v[0] *= r
	v[1] *= r
	v[2] *= r
}

// Normalize sets the vector the normalized value.
func (v *Vector3) Normalize() {
	m := v.Magnitude()
	if !FloatsEqual(m, 0.0) {
		l := 1.0 / m
		v[0] *= l
		v[1] *= l
		v[2] *= l
	}
}

// Set sets the vector equal to the values of the second vector.
func (v *Vector3) Set(v2 *Vector3) {
	v[0] = v2[0]
	v[1] = v2[1]
	v[2] = v2[2]
}

// Sub subtracts a second vector from this vector.
func (v *Vector3) Sub(v2 *Vector3) {
	v[0] -= v2[0]
	v[1] -= v2[1]
	v[2] -= v2[2]
}

// MulWith multiplies a vector by a float64 number.
func (v *Vector4) MulWith(r float64) {
	v[0] *= r
	v[1] *= r
	v[2] *= r
	v[3] *= r
}
