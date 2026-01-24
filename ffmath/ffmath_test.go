// SPDX-FileCopyrightText: 2026 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package ffmath

import (
	"testing"

	"github.com/orsinium-labs/tinymath"
)

func TestEqualApprox(t *testing.T) {
	tests := []struct {
		name string
		a, b float32
		want bool
	}{
		{name: "inf=inf", a: tinymath.Inf, b: tinymath.Inf, want: true},
		{name: "neginf=neginf", a: tinymath.NegInf, b: tinymath.NegInf, want: true},
		{name: "inf!=neginf", a: tinymath.Inf, b: tinymath.NegInf, want: false},
		{name: "neginf!=inf", a: tinymath.NegInf, b: tinymath.Inf, want: false},
		{name: "integers equal", a: 105, b: 105, want: true},
		{name: "integers not equal", a: 105, b: 106, want: false},
		{name: "fraction equal", a: 1.005, b: 1.005, want: true},
		{name: "fraction not equal", a: 1.005, b: 1.006, want: false},
		{name: "tiny fraction equal", a: 1.000001, b: 1.000006, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := EqualApprox(test.a, test.b)
			if got != test.want {
				t.Errorf("EqualApprox(%v, %v): want %t, got %t", test.a, test.b, test.want, got)
			}
		})
	}
}
