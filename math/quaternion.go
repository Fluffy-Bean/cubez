// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package math

import (
	"math"
)

// Quaternion is an alias for Vector4
type Quaternion Vector4

// QuatFromAxis creates an quaternion from an axis and an angle.
func QuatFromAxis(angle, x, y, z float64) Quaternion {
	s := math.Sin(angle / 2.0)
	c := math.Cos(angle / 2.0)

	result := Quaternion{c, x * s, y * s, z * s}
	result.Normalize()
	return result
}

// AddScaledVector adds the given vector to this quaternion, scaled
// by the given amount.
func (q *Quaternion) AddScaledVector(v *Vector3, scale float64) {
	var temp Quaternion

	temp.X = v.X * scale
	temp.Y = v.Y * scale
	temp.Z = v.Z * scale
	temp.W = 0.0

	temp.Multiply(q)

	q.X += temp.X * 0.5
	q.Y += temp.Y * 0.5
	q.Z += temp.Z * 0.5
	q.W += temp.W * 0.5
}

// SetIdentity loads the quaternion with its identity.
func (q *Quaternion) SetIdentity() {
	q.X, q.Y, q.Z, q.W = 0.0, 0.0, 0.0, 1.0
}

// Len returns the length of the quaternion.
func (q *Quaternion) Len() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

// Multiply multiplies the quaternion by another quaternion.
func (q *Quaternion) Multiply(q2 *Quaternion) {
	var w, x, y, z float64
	w = q.W*q2.W - q.X*q2.X - q.Y*q2.Y - q.Z*q2.Z
	x = q.W*q2.X + q.X*q2.W + q.Y*q2.Z - q.Z*q2.Y
	y = q.W*q2.Y + q.Y*q2.W + q.Z*q2.X - q.X*q2.Z
	z = q.W*q2.Z + q.Z*q2.W + q.X*q2.Y - q.Y*q2.X
	q.W, q.X, q.Y, q.Z = w, x, y, z
}

// Rotate rotates a vector by the rotation this quaternion represents.
func (q *Quaternion) Rotate(v *Vector3) Vector3 {
	qvec := Vector3{q.X, q.Y, q.Z}
	cross := qvec.Cross(v)

	// v + 2q_w * (q_v x v) + 2q_v x (q_v x v)
	result := *v

	qvec.Scale(2.0)
	c2 := qvec.Cross(&cross)
	result.Add(&c2)

	cross.Scale(2.0 * q.W)
	result.Add(&cross)
	return result
}

// Conjugated returns the conjugate of a quaternion. Equivalent to Quaternion{w,-x,-y,-z}.
func (q *Quaternion) Conjugated() Quaternion {
	return Quaternion{-q.X, -q.Y, -q.Z, q.W}
}

// Normalize normalizes the quaternion.
func (q *Quaternion) Normalize() {
	length := q.Len()

	if FloatsEqual(1.0, length) {
		return
	}

	if length == 0.0 {
		q.SetIdentity()
		return
	}

	if length == InfPos {
		length = MaxValue
	}

	invLength := 1.0 / length
	q.Scale(invLength)
}

// LookAt sets the quaternion to the orientation needed to look at a 'center' from
// the 'eye' position with 'up' as a reference vector for the up direction.
// Note: this was modified from the go-gl/mathgl library.
func (q *Quaternion) LookAt(eye, center, up *Vector3) {
	direction := center
	direction.Subtract(eye)
	direction.Normalize()

	// Find the rotation between the front of the object (that we assume towards Z-,
	// but this depends on your model) and the desired direction
	rotDir := QuatBetweenVectors(&Vector3{0, 0, -1}, direction)

	// Recompute up so that it's perpendicular to the direction
	// You can skip that part if you really want to force up
	//right := direction.Cross(up)
	//up = right.Cross(direction)

	// Because of the 1rst rotation, the up is probably completely screwed up.
	// Find the rotation between the "up" of the rotated object, and the desired up
	upCur := rotDir.Rotate(&Vector3{0, 1, 0})
	rotTarget := QuatBetweenVectors(&upCur, up)

	rotTarget.Multiply(&rotDir) // remember, in reverse order.
	rotTarget.Inverse()         // camera rotation should be inversed!

	q.X = rotTarget.X
	q.Y = rotTarget.Y
	q.Z = rotTarget.Z
	q.W = rotTarget.W
}

// QuatBetweenVectors calculates the rotation between two vectors.
// Note: this was modified from the go-gl/mathgl library.
func QuatBetweenVectors(s, d *Vector3) Quaternion {
	start := *s
	dest := *d
	start.Normalize()
	dest.Normalize()

	cosTheta := start.Dot(&dest)
	if cosTheta < -1.0+Epsilon {
		// special case when vectors in opposite directions:
		// there is no "ideal" rotation axis
		// So guess one; any will do as long as it's perpendicular to start
		posX := Vector3{1.0, 0.0, 0.0}
		axis := posX.Cross(&start)
		if axis.Dot(&axis) < Epsilon {
			// bad luck, they were parallel, try again!
			posY := Vector3{0.0, 1.0, 0.0}
			axis = posY.Cross(&start)
		}

		axis.Normalize()
		return QuatFromAxis(math.Pi, axis.X, axis.Y, axis.Z)
	}

	axis := start.Cross(&dest)
	ang := math.Sqrt((1.0 + cosTheta) * 2.0)
	axis.Scale(1.0 / ang)

	return Quaternion{
		ang * 0.5,
		axis.X, axis.Y, axis.Z,
	}
}

// Inverse calculates the inverse of a quaternion. The inverse is equivalent
// to the conjugate divided by the square of the length.
//
// This method computes the square norm by directly adding the sum
// of the squares of all terms instead of actually squaring q1.Len(),
// both for performance and percision.
func (q *Quaternion) Inverse() {
	c := q.Conjugated()
	c.Scale(1.0 / q.Dot(q))
	q.X = c.X
	q.Y = c.Y
	q.Z = c.Z
	q.W = c.W
}

// Scale scales every element of the quaternion by some constant factor.
func (q *Quaternion) Scale(c float64) {
	q.X *= c
	q.Y *= c
	q.Z *= c
	q.W *= c
}

// Dot calculates the dot product between two quaternions, equivalent to if this was a Vector4
func (q *Quaternion) Dot(q2 *Quaternion) float64 {
	return q.Y*q2.Y + q.Y*q2.Y + q.Z*q2.Z + q.W*q2.W
}
