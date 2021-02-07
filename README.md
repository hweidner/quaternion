[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDocs](https://godocs.io/github.com/hweidner/quaternion?status.svg)](https://godocs.io/github.com/hweidner/quaternion)
[![Go Reference](https://pkg.go.dev/badge/github.com/hweidner/quaternion.svg)](https://pkg.go.dev/github.com/hweidner/quaternion)
[![Go Report Card](https://goreportcard.com/badge/github.com/hweidner/quaternion)](https://goreportcard.com/report/github.com/hweidner/quaternion)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/hweidner/quaternion.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/hweidner/quaternion/alerts/)

quaternion
==========

The quaternions are a 4-tuple number system and an extension to complex
numbers. See http://en.wikipedia.org/wiki/Quaternion for a definition and
description of their properties.

This packages provides a type for storing a quaternion and basic calculus
operations for Go.

Examples
--------

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

	// Calculate the product of two quaterions
	prod := x.Mult(y)

Rationale
---------

All operations are invoked in an object oriented style. They return the
result (either as type Quaternion or float64).

Unary operations have the form x.Op(). The value x is not changed.
Binary operations are called as a x.Op(y) where x is the left and
y the right operand. The values of x and y are not changed. All methods
return a value of type float64 or Quaternion.

TODOs
-----

This package is in an early stage. More functions need to be implemented, e.g.
power, trigonometry functions.

Copyright and License
---------------------

This package is released under the MIT license.
The full license text can be found in the [LICENSE](LICENSE) file.
