/**
signal.go
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
)

type DiscreteSignal interface {
	Get(N int) complex128
	Range() (Min int, Max int)
	MaxAmplitude() float64
}

type Delta struct {
	Offset    int
	Amplitude complex128
}

func (d *Delta) Get(N int) complex128 {
	if N == d.Offset {
		return d.Amplitude
	}
	return 0
}

func (d *Delta) Range() (Min int, Max int) {
	return d.Offset, d.Offset
}

func (d *Delta) MaxAmplitude() float64 {
	return cmplx.Abs(d.Amplitude)
}

var DiracDelta Delta = Delta{0, complex128(1.0)}

type Step struct {
	Offset    int
	Amplitude complex128
}

func (step *Step) Get(N int) complex128 {
	if N >= step.Offset {
		return step.Amplitude
	}
	return 0.0
}

func (step *Step) Range() (Min int, Max int) {
	return step.Offset, int(^uint(0) >> 1) // AKA MaxInt
}

func (step *Step) MaxAmplitude() float64 {
	return cmplx.Abs(step.Amplitude)
}

var UnitStep = Step{0, complex128(1.0)}
