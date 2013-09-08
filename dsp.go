/**
 * Created with IntelliJ IDEA.
 * User: nick
 * Date: 8/31/13
 * Time: 11:10 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"log"
	"math"
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

func drawGraph(drawingArea *gtk.DrawingArea, pixMap *gdk.Pixmap, gc *gdk.GC) {
	allocation := drawingArea.GetAllocation()
	drawable := pixMap.GetDrawable()
	scaleY := func(val float64) int {
		return allocation.Height - int(5.0+float32(val+3.0)*float32(allocation.Height-10)/6.0)
	}
	scaleX := func(val float64) int {
		return int(5.0 + float32(val+10.0)*float32(allocation.Width-10)/20.0)
	}
	gc.SetRgbFgColor(gdk.NewColor("white"))
	drawable.DrawRectangle(gc, true, 0, 0, -1, -1)
	gc.SetRgbBgColor(gdk.NewColor("white"))
	gc.SetRgbFgColor(gdk.NewColor("grey"))
	/* Going to draw a unit grid, +- 10 horiz nad += 3 verticallu */
	for y := float64(-2); y <= float64(2); y++ {
		drawable.DrawLine(gc, 5, scaleY(y), allocation.Width-5, scaleY(y))
	}
	for x := float64(-9); x <= float64(9); x++ {
		drawable.DrawLine(gc, scaleX(x), 5, scaleX(x), allocation.Height-5)
	}
	gc.SetRgbFgColor(gdk.NewColor("black"))
	drawable.DrawLine(gc, 5, scaleY(0.0), allocation.Width-5, scaleY(0.0))
	drawable.DrawLine(gc, scaleX(0.0), 5, scaleX(0.0), allocation.Height-5)
	gc.SetRgbFgColor(gdk.NewColor("blue"))
	for x := -math.Pi * 3.0; x <= math.Pi*3.0; x += 0.1 {
		drawable.DrawLine(gc, scaleX(x), scaleY(0), scaleX(x), scaleY(3.0*math.Sin(x)/x))
	}
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("DSP Playground")

	window.Connect("delete-event", deleteEvent)
	window.Connect("destroy", destroyEvent)
	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)

	drawingArea := gtk.NewDrawingArea()
	var pixmap *gdk.Pixmap
	var gc *gdk.GC
	drawingArea.Connect("expose_event", func() {
		if pixmap != nil {
			drawingArea.GetWindow().GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
		}
	})
	drawingArea.Connect("configure_event", func(ctx *glib.CallbackContext) {
		if pixmap != nil {
			pixmap.Unref()
		}
		allocation := drawingArea.GetAllocation()
		pixmap = gdk.NewPixmap(drawingArea.GetWindow().GetDrawable(),
			allocation.Width,
			allocation.Height,
			-1)
		gc = gdk.NewGC(pixmap.GetDrawable())
		drawGraph(drawingArea, pixmap, gc)
	})

	drawingArea.SetEvents(int(gdk.EXPOSE | gdk.CONFIGURE))
	vbox.Add(drawingArea)
	window.Add(vbox)
	window.SetSizeRequest(1010, 300)
	window.ShowAll()

	gtk.Main()
}
