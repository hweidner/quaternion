// Copyright 2013 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the GNU Lesser General Public License
// Version 3 that can be found in the LICENSE file.

/*
Package quaternion implements the quaternions, a 4-tuple number system.

The quaternions are a 4-tuple number system and an extension to complex
numbers. See http://en.wikipedia.org/wiki/Quaternion for a definition and
description of their properties.

This packages provides a type for storing a quaternion and basic calculus
operations.

Examples

Some examples of quaternion usage:

	import "quaternion"

	// Construction by tuple
	x := quaternion.Quaternion{1, 2, -3, -1}

	// Construction from real number by function
	y := quaternion.FromFloat64(0.75)

	// Calculate the inverse of a quaternion
	inv := x.Inv()

	// Calculate the sum of two quaternions
	sum := x.Add(y)

	// Calcualte the product of two quaterions
	prod := x.Mult(y)

Rationale

All operations are invoked in an object oriented style. They return the
result (either as type Quaternion or float64).

Unary operations have the form x.Op(). The value x is not changed.
Binary operations are called as a x.Op(y) where x is the left and
y the right operand. The values of x and y are not changed. All methods
return a value of type float64 or Quaternion.

TODOs

This package is in an early stage. More functions need to be implemented, e.g.
power, trigonometry functions.

License

This package is released under the GNU Lesser General Public License, Version
3. The full license text can be found in the LICENSE file of the source code
distribution.

*/
package quaternion
