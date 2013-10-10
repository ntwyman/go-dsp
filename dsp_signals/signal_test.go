/**
singal_test.go
Copyright (c) 2013 Nick Twyman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. To change this template use File | Settings | File Templates.
*/
package signal

import (
	"math/cmplx"
	"testing"
)

func doDeltaTest(t *testing.T, delta DiscreteSignal, offset int, scale complex128) {

	if min, max := delta.Range(); min != offset || max != offset {
		t.Error("Range for delta")
	}

	if delta.MaxAmplitude() != cmplx.Abs(scale) {
		t.Error("Amplitude for delta")
	}

	for i := -1000; i < 1000; i++ {
		if val := delta.Get(i); i == offset && val != scale {
			t.Error("Delta value")
		} else if i != offset && val != 0.0 {
			t.Error("Delta zero value")
		}
	}
}

func TestDeltaDirac(t *testing.T) {

	doDeltaTest(t, &DiracDelta, 0, 1.0)
}

func TestDelta(t *testing.T) {
	delta := Delta{13, complex128(2 + 3i)}
	doDeltaTest(t, &delta, 13, 2+3i)
	minusDelta := Delta{-23, complex128(-3 + 2i)}
	doDeltaTest(t, &minusDelta, -23, -3+2i)
}

func doStepTest(t *testing.T, step DiscreteSignal, offset int, scale complex128) {
	if min, max := step.Range(); min != offset || max < int(^uint(0)>>1) {
		t.Error("Range for step function")
	}

	if step.MaxAmplitude() != cmplx.Abs(scale) {
		t.Error("Amplitude for step")
	}

	for i := -1000; i < 1000; i++ {
		if val := step.Get(i); i >= offset && val != scale {
			t.Error("Step above offset")
		} else if i < offset && val != 0.0 {
			t.Error("Step below offset")
		}
	}
}

func TestUnitStep(t *testing.T) {
	doStepTest(t, &UnitStep, 0, complex128(1.0))
}

func TestStep(t *testing.T) {
	step := Step{534, complex128(3.14 + 0.0596i)}
	doStepTest(t, &step, 534, complex128(3.14+0.0596i))
	minusStep := Step{-435, complex128(-12345i)}
	doStepTest(t, &minusStep, -435, complex128(-12345i))

}
