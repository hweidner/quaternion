// Copyright 2013 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the MIT license. See the LICENSE file
// for a full text of the license.
// SPDX-License-Identifier: MIT

package quaternion

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func almostEqualQ(a, b Quaternion) bool {
	return math.Abs(a.Qr-b.Qr) < 1e-6 &&
		math.Abs(a.Qi-b.Qi) < 1e-6 &&
		math.Abs(a.Qj-b.Qj) < 1e-6 &&
		math.Abs(a.Qk-b.Qk) < 1e-6
}

func TestQuaternion(t *testing.T) {
	// do tests and report failure
	doTest := func(cond bool, msg string) {
		if !cond {
			t.Error(msg)
		}
	}

	// some quaternions
	x := Quaternion{1, 2, -0.5, -1}
	y := Quaternion{-3, 4, 0, 2}
	z := Quaternion{-2, 1, 2, -4}

	// test specific calculation results
	doTest(almostEqual(y.Norm(), 29.0), "Norm() test failed.")
	doTest(almostEqual(z.Abs(), 5.0), "Abs() test failed.")
	doTest(almostEqualQ(z.Signum(), Quaternion{-0.4, 0.2, 0.4, -0.8}),
		"Signum() Test failed.")
	doTest(almostEqualQ(x.Exp().Log(), x), "Exp()/Log() Test failed.")
	doTest(almostEqualQ(x.Mult(y), Quaternion{-9.0, -3.0, -6.5, 7.0}),
		"Mult() test failed.")
	doTest(almostEqual(y.ScalarProd(z), 2.0), "ScalarProd() test failed.")
	doTest(almostEqualQ(z.CrossProd(y), Quaternion{0, 4, -18, -8}),
		"CrossProd() test failed")

	// test formulae that hold for each quaternion
	for _, q := range [...]Quaternion{x, y, z} {
		doTest(almostEqualQ(q.Add(q.Conj()), FromFloat64(2*q.Real())),
			"Conj()/Real() test failed.")
		doTest(almostEqual(q.Signum().Norm(), 1.0), "Signum()/Norm() test failed")
		doTest(almostEqualQ(q.Sub(FromFloat64(q.Real())).Sub(q.Img()),
			Null()), "Real()/Img()/Sub() test failed.")
		doTest(almostEqualQ(q.Inv().Mult(q), One()),
			"Inv()/Mult() test failed.")
		doTest(almostEqualQ(q.Mult(q.Inv()), One()),
			"Mult()/Inv() test failed.")
		doTest(almostEqualQ(q.Log().Exp(), q),
			"Log()/Exp() test failed.")
	}
}
