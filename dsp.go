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
