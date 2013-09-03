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
	"log"
	"os"
)

func deleteEvent(ctx *glib.CallbackContext) bool {
	log.Println("delete-event called", ctx)
	return true
}

func destroyEvent(ctx *glib.CallbackContext) {
	log.Println("destroy-event called", ctx)
	gtk.MainQuit()
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

	window.Connect("delete-event", deleteEvent)
	window.Connect("destroy", destroyEvent)
	window.SetBorderWidth(10)

	button := gtk.NewButtonWithLabel("Hello World!")
	button.Clicked(func() {
		log.Println("Hello world!")
		window.Destroy()
	})

	window.Add(button)
	button.Show()
	window.Show()

	gtk.Main()
}
