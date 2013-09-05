/**
 * Created with IntelliJ IDEA.
 * User: nick
 * Date: 8/31/13
 * Time: 11:10 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdk"
	"log"
	"os"
)

func deleteEvent(ctx *glib.CallbackContext) bool {
	log.Println("delete-event called", ctx)
	return false
}

func destroyEvent(ctx *glib.CallbackContext) {
	log.Println("destroy-event called", ctx)
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
		gc.SetRgbFgColor(gdk.NewColor("#c0dcc0"))
		pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)
		gc.SetRgbBgColor(gdk.NewColor("#c0dcc0"))
		gc.SetRgbFgColor(gdk.NewColor("black"))
	})

	drawingArea.SetEvents(int(gdk.EXPOSE | gdk.CONFIGURE))
	vbox.Add(drawingArea)
	window.Add(vbox)
	window.SetSizeRequest(400, 300)
	window.ShowAll()

	gtk.Main()
}
