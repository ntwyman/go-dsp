/**
graph.go
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

const graphMargin float64 = 5.0

type Graph struct {
	gtk.DrawingArea
	pixmap         *gdk.Pixmap
	gc             *gdk.GC
	minX           float64
	maxX           float64
	minY           float64
	maxY           float64
	xMinorInterval float64
	xMajorInterval float64
	yMinorInterval float64
	yMajorInterval float64
	scaleX         func(float64) int
	scaleY         func(float64) int
}

func NewGraph() *Graph {
	graph := &Graph{*gtk.NewDrawingArea(), nil, nil, -10.0, 10.0, -4.0, 4.0, 0.5, 100.0, 0.5, 100.0, nil, nil}
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
		graph.setScale()
		graph.plot()
	})

	graph.SetEvents(int(gdk.EXPOSE | gdk.CONFIGURE))
	return graph
}

func (g *Graph) setScale() {
	allocation := g.GetAllocation()
	g.scaleY = func(val float64) int {
		return allocation.Height - int(graphMargin+(val-g.minY)*(float64(allocation.Height)-graphMargin*2)/(g.maxY-g.minY))
	}
	g.scaleX = func(val float64) int {
		return int(graphMargin + ((val-g.minX)*(float64(allocation.Width)-graphMargin*2))/(g.maxX-g.minX))
	}
}

func (g *Graph) clear() {
	g.gc.SetRgbFgColor(gdk.NewColor("white"))
	g.pixmap.GetDrawable().DrawRectangle(g.gc, true, 0, 0, -1, -1)
	g.gc.SetRgbBgColor(gdk.NewColor("white"))
}

func (g *Graph) drawLineScaled(x1 float64, y1 float64, x2 float64, y2 float64) {
	g.pixmap.GetDrawable().DrawLine(g.gc, g.scaleX(x1), g.scaleY(y1), g.scaleX(x2), g.scaleY(y2))
}

func (g *Graph) drawGrid(xInterval float64, yInterval float64) {

	if xInterval > 0 {
		for x := math.Ceil((g.minX+(g.maxX-g.minX)/1000)/xInterval) * xInterval; x < g.maxX; x += xInterval {
			g.drawLineScaled(x, g.minY, x, g.maxY)
		}
	}
	if yInterval > 0 {
		for y := math.Ceil((g.minY+(g.maxY-g.minY)/1000)/yInterval) * yInterval; y < g.maxY; y += yInterval {
			g.drawLineScaled(g.minX, y, g.maxX, y)
		}
	}
}

func (g *Graph) plotGrid() {

	// First draw the minor grid
	g.gc.SetRgbFgColor(gdk.NewColor("grey"))
	g.drawGrid(g.xMinorInterval, g.yMinorInterval)

	// And the major grid
	g.gc.SetRgbFgColor(gdk.NewColor("black"))
	g.drawGrid(g.xMajorInterval, g.yMajorInterval)
}

func (g *Graph) plot() {
	g.clear()
	g.plotGrid()

	g.gc.SetRgbFgColor(gdk.NewColor("blue"))
	for x := -math.Pi * 3.0; x <= math.Pi*3.0; x += 0.1 {
		g.drawLineScaled(x, 0, x, 3.0*math.Sin(x)/x)
	}
}
