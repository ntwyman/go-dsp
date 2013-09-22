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
	"testing"
)

func TestDeltaDirac(t *testing.T) {

	if min, max := DiracDelta.Range(); min != 0 || max != 0 {
		t.Error("Range for DiracDel")
	}

	if DiracDelta.MaxAmplitude() != 1.0 {
		t.Error("Amplitude for dirac delta")
	}

	for i := -1000; i < 1000; i++ {
		if val := DiracDelta.Get(i); i == 0 && val != 1.0 {
			t.Error("DetlaDirec zero value")
		} else if i != 0 && val != 0.0 {
			t.Error("DetlaDirec non-zero value")
		}
	}
}
