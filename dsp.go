/**
dsp.go
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
package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/ntwyman/go-dsp/dsp_gtk"
	"os"
)

func deleteEvent(ctx *glib.CallbackContext) bool {
	//log.Println("delete-event called", ctx)
	return false
}

func destroyEvent(ctx *glib.CallbackContext) {
	//log.Println("destroy-event called", ctx)
	gtk.MainQuit()
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("DSP Playground")

	window.Connect("delete-event", deleteEvent)
	window.Connect("destroy", destroyEvent)
	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)

	vbox.Add(dsp_gtk.NewGraph())
	window.Add(vbox)
	window.SetSizeRequest(1010, 300)
	window.ShowAll()

	gtk.Main()
}
