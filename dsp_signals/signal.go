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
	offset    int
	amplitude complex128
}

func (d *Delta) Get(N int) complex128 {
	if N == d.offset {
		return d.amplitude
	}
	return 0
}

func (d *Delta) Range() (Min int, Max int) {
	return d.offset, d.offset
}

func (d *Delta) MaxAmplitude() float64 {
	return cmplx.Abs(d.amplitude)
}

var DiracDelta Delta = Delta{0, complex128(1.0)}
