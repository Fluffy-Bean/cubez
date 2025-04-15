// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package math

import (
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

// Vector4 is a vector of four floats.
type Vector4 [4]float64

// Quat is the type of a quaternion (w,x,y,z).
type Quat Vector4

// Clear sets the vector to {0.0, 0.0, 0.0}.
func (v *Vector3) Clear() {
	v.X, v.Y, v.Z = 0.0, 0.0, 0.0
}

// GetIndex is a helper method for getting XYZ vales by... an index
func (v *Vector3) GetIndex(i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	}

	panic("Vector3 index out of range")
}

// SetIndex is a helper method for setting XYZ vales by... an index
func (v *Vector3) SetIndex(i int, value float64) {
	switch i {
	case 0:
		v.X = value
		return
	case 1:
		v.Y = value
		return
	case 2:
		v.Z = value
		return
	}

	panic("Vector3 index out of range")
}

// Add adds a vector to another vector.
func (v *Vector3) Add(v2 *Vector3) {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
}

// Add adds a vector, scaled by a float64, to another vector.
func (v *Vector3) AddScaled(v2 *Vector3, scale float64) {
	v.X += v2.X * scale
	v.Y += v2.Y * scale
	v.Z += v2.Z * scale
}

// ComponentProduct performs a component-wise product with another vector.
func (v *Vector3) ComponentProduct(v2 *Vector3) {
	v.X *= v2.X
	v.Y *= v2.Y
	v.Z *= v2.Z
}

// Cross returns the cross product of this vector with another.
func (v *Vector3) Cross(v2 *Vector3) Vector3 {
	return Vector3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

// Dot returns the dot product of this vector with another.
func (v *Vector3) Dot(v2 *Vector3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Magnitude returns the magnitude of the vector.
func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// SquareMagnitude returns the magitude of the vector, squared.
func (v *Vector3) SquareMagnitude() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Scale multiplies a vector by a float64 number.
func (v *Vector3) Scale(scale float64) {
	v.X *= scale
	v.Y *= scale
	v.Z *= scale
}

// Normalize sets the vector the normalized value.
func (v *Vector3) Normalize() {
	m := v.Magnitude()
	if !FloatsEqual(m, 0.0) {
		l := 1.0 / m
		v.X *= l
		v.Y *= l
		v.Z *= l
	}
}

// Set sets the vector equal to the values of the second vector.
func (v *Vector3) Set(v2 *Vector3) {
	v.X = v2.X
	v.Y = v2.Y
	v.Z = v2.Z
}

// Subtract subtracts a second vector from this vector.
func (v *Vector3) Subtract(v2 *Vector3) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}

// Scale multiplies a vector by a float64 number.
func (v *Vector4) Scale(scale float64) {
	v[0] *= scale
	v[1] *= scale
	v[2] *= scale
	v[3] *= scale
}
