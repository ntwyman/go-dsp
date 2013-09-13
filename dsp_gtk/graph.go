/**
gtk/graph.go
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
package dsp_gtk

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"math"
)

type Graph struct {
	gtk.DrawingArea
	pixmap *gdk.Pixmap
	gc     *gdk.GC
	minX   float64
	maxX   float64
	minY   float64
	maxY   float64
}

func NewGraph() *Graph {
	graph := &Graph{*gtk.NewDrawingArea(), nil, nil, -1, 1, -1, 1}
	graph.Connect("expose_event", func() {
		if graph.pixmap != nil {
			graph.GetWindow().GetDrawable().DrawDrawable(graph.gc, graph.pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
		}
	})
	graph.Connect("configure_event", func(ctx *glib.CallbackContext) {
		if graph.pixmap != nil {
			graph.pixmap.Unref()
		}
		allocation := graph.GetAllocation()
		graph.pixmap = gdk.NewPixmap(graph.GetWindow().GetDrawable(),
			allocation.Width,
			allocation.Height,
			-1)
		graph.gc = gdk.NewGC(graph.pixmap.GetDrawable())
		graph.Plot()
	})
	graph.SetEvents(int(gdk.EXPOSE | gdk.CONFIGURE))
	return graph
}

func (g *Graph) Plot() {
	allocation := g.GetAllocation()
	drawable := g.pixmap.GetDrawable()
	scaleY := func(val float64) int {
		return allocation.Height - int(5.0+float32(val+3.0)*float32(allocation.Height-10)/6.0)
	}
	scaleX := func(val float64) int {
		return int(5.0 + float32(val+10.0)*float32(allocation.Width-10)/20.0)
	}
	g.gc.SetRgbFgColor(gdk.NewColor("white"))
	drawable.DrawRectangle(g.gc, true, 0, 0, -1, -1)
	g.gc.SetRgbBgColor(gdk.NewColor("white"))
	g.gc.SetRgbFgColor(gdk.NewColor("grey"))
	/* Going to draw a unit grid, +- 10 horiz nad += 3 verticallu */
	for y := float64(-2); y <= float64(2); y++ {
		drawable.DrawLine(g.gc, 5, scaleY(y), allocation.Width-5, scaleY(y))
	}
	for x := float64(-9); x <= float64(9); x++ {
		drawable.DrawLine(g.gc, scaleX(x), 5, scaleX(x), allocation.Height-5)
	}
	g.gc.SetRgbFgColor(gdk.NewColor("black"))
	drawable.DrawLine(g.gc, 5, scaleY(0.0), allocation.Width-5, scaleY(0.0))
	drawable.DrawLine(g.gc, scaleX(0.0), 5, scaleX(0.0), allocation.Height-5)
	g.gc.SetRgbFgColor(gdk.NewColor("blue"))
	for x := -math.Pi * 3.0; x <= math.Pi*3.0; x += 0.1 {
		drawable.DrawLine(g.gc, scaleX(x), scaleY(0), scaleX(x), scaleY(3.0*math.Sin(x)/x))
	}
}
