package main

import (
	"github.com/niemeyer/qml"
	"github.com/martin-angelov1992/GoCarousel"
	"fmt"
)

func main() {
	fmt.Println("start")
	qml.Init(nil)
	engine := qml.NewEngine()
	base, _ := engine.LoadFile("base.qml")
	win := base.CreateWindow(nil)
	carousel := GoCarousel.NewCarousel(engine, win)
	carousel.SetImages([]string{"../test_images/ubuntu-gopher.png", "../test_images/firefox.png", "../test_images/fire.jpg"})
	win.Show()
	carousel.Run()
	win.Wait()
	fmt.Println("end")
}
