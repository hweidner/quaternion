// Copyright 2013 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the MIT license. See the LICENSE file
// for a full text of the license.
// SPDX-License-Identifier: MIT

package quaternion

import "math"

// VERSION is the version number of package quaternion.
const VERSION = "0.1"

// Quaternion stores a quaternion as a 4-tuple of float64.
type Quaternion struct {
	Qr, Qi, Qj, Qk float64
}

// fixed quaternions for internal use.
var (
	q0 = Quaternion{0, 0, 0, 0}
	q1 = Quaternion{1, 0, 0, 0}
	qi = Quaternion{0, 1, 0, 0}
	qj = Quaternion{0, 0, 1, 0}
	qk = Quaternion{0, 0, 0, 1}
)

// Constructors

// Null returns the empty quaternion.
func Null() Quaternion {
	return q0
}

// One returns the real quaternion which equals one.
func One() Quaternion {
	return q1
}

// I returns the quaternion which equals i.
func I() Quaternion {
	return qi
}

// J returns the quaternion which equals j.
func J() Quaternion {
	return qj
}

// K returns the quaternion which equals k.
func K() Quaternion {
	return qk
}

// FromFloat64 converts a float64 value into a real quaternion.
func FromFloat64(a float64) (b Quaternion) {
	b.Qr = a
	return
}

// Unary operations

// Conj calculates the conjugation of a quaterion.
func (a Quaternion) Conj() (b Quaternion) {
	b.Qr = a.Qr
	b.Qi = -a.Qi
	b.Qj = -a.Qj
	b.Qk = -a.Qk
	return
}

// Norm calculates the (quadratic) norm of a quaternion.
func (a Quaternion) Norm() float64 {
	return a.Qr*a.Qr + a.Qi*a.Qi + a.Qj*a.Qj + a.Qk*a.Qk
}

// Abs calculates the absolute value of a quaternion
// (sometimes referred to as length, norm, magnitude, or modulus).
func (a Quaternion) Abs() float64 {
	return math.Sqrt(a.Qr*a.Qr + a.Qi*a.Qi + a.Qj*a.Qj + a.Qk*a.Qk)
}

// Real returns the real part of a quaternion as a float64 value.
func (a Quaternion) Real() float64 {
	return a.Qr
}

// Img returns the imaginary (vector) part of a quaternion.
func (a Quaternion) Img() (b Quaternion) {
	b.Qi = a.Qi
	b.Qj = a.Qj
	b.Qk = a.Qk
	return
}

// Signum returns the signum of a quaternion. Unless zero, it always has the
// norm of 1.
func (a Quaternion) Signum() (b Quaternion) {
	norm := a.Qr*a.Qr + a.Qi*a.Qi + a.Qj*a.Qj + a.Qk*a.Qk

	if norm == 0.0 {
		return
	}

	abs := math.Sqrt(norm)
	b.Qr = a.Qr / abs
	b.Qi = a.Qi / abs
	b.Qj = a.Qj / abs
	b.Qk = a.Qk / abs
	return
}

// Inv calculates the inverse of a quaternion. The inverse a' to a quaternion
// a is the quaternion for which the formulae a * a' = 1 and a' * a = 1 are true.
func (a Quaternion) Inv() (b Quaternion) {
	norm := a.Qr*a.Qr + a.Qi*a.Qi + a.Qj*a.Qj + a.Qk*a.Qk

	if norm == 0.0 {
		return
	}

	b.Qr = a.Qr / norm
	b.Qi = -a.Qi / norm
	b.Qj = -a.Qj / norm
	b.Qk = -a.Qk / norm
	return
}

// Exp calculates the exponentation of a quaternion, i.e. e^q where e is the
// Euler number.
func (a Quaternion) Exp() (b Quaternion) {
	exp := math.Exp(a.Qr)
	unit := math.Sqrt(a.Qi*a.Qi + a.Qj*a.Qj + a.Qk*a.Qk)

	if unit == 0.0 {
		b.Qr = exp
		return
	}

	real := math.Cos(unit)
	imgfact := exp * math.Sin(unit) / unit

	b.Qr = exp * real
	b.Qi = imgfact * a.Qi
	b.Qj = imgfact * a.Qj
	b.Qk = imgfact * a.Qk
	return
}

// Log calculates the natural logarithm of a quaternion.
func (a Quaternion) Log() (b Quaternion) {
	imgnorm := a.Qi*a.Qi + a.Qj*a.Qj + a.Qk*a.Qk
	abs := math.Sqrt(a.Qr*a.Qr + imgnorm)
	unit := math.Sqrt(imgnorm)

	if unit == 0.0 {
		b.Qr = math.Log(abs)
		return
	}

	imgfact := math.Acos(a.Qr/abs) / unit

	b.Qr = math.Log(abs)
	b.Qi = imgfact * a.Qi
	b.Qj = imgfact * a.Qj
	b.Qk = imgfact * a.Qk
	return
}

// Binary operations

// Add calculates the sum of two quaternions, which is the component-wise
// sum of the four tuples.
func (a Quaternion) Add(b Quaternion) (c Quaternion) {
	c.Qr = a.Qr + b.Qr
	c.Qi = a.Qi + b.Qi
	c.Qj = a.Qj + b.Qj
	c.Qk = a.Qk + b.Qk
	return
}

// Sub calculates the difference (subtraction) of two quaternions, which is the
// component-wise difference of the four tuples.
func (a Quaternion) Sub(b Quaternion) (c Quaternion) {
	c.Qr = a.Qr - b.Qr
	c.Qi = a.Qi - b.Qi
	c.Qj = a.Qj - b.Qj
	c.Qk = a.Qk - b.Qk
	return
}

// Mult calculates the product of two quaternions. Note that the
// multiplication is usually not commutative, e.g. a.Mult(b) != b.Mult(a).
func (a Quaternion) Mult(b Quaternion) (c Quaternion) {
	c.Qr = a.Qr*b.Qr - a.Qi*b.Qi - a.Qj*b.Qj - a.Qk*b.Qk
	c.Qi = a.Qr*b.Qi + a.Qi*b.Qr + a.Qj*b.Qk - a.Qk*b.Qj
	c.Qj = a.Qr*b.Qj - a.Qi*b.Qk + a.Qj*b.Qr + a.Qk*b.Qi
	c.Qk = a.Qr*b.Qk + a.Qi*b.Qj - a.Qj*b.Qi + a.Qk*b.Qr
	return
}

// ScalarProd calculates the scalar product of two quaternions, which
// is a real number.
func (a Quaternion) ScalarProd(b Quaternion) float64 {
	return a.Qr*b.Qr + a.Qi*b.Qi + a.Qj*b.Qj + a.Qk*b.Qk
}

// CrossProd calculates the cross product of two quaternions.
func (a Quaternion) CrossProd(b Quaternion) (c Quaternion) {
	c.Qr = 0
	c.Qi = a.Qj*b.Qk - a.Qk*b.Qj
	c.Qj = -a.Qi*b.Qk + a.Qk*b.Qi
	c.Qk = a.Qi*b.Qj - a.Qj*b.Qi
	return
}
