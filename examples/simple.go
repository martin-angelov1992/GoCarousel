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
	err, carousel := GoCarousel.NewCarousel(engine, win)
	if err != nil {
		fmt.Println(err)
		return
	}
	carousel.SetImages([]string{"ubuntu-gopher.png", "firefox.png"})
	win.Show()
	win.Wait()
	fmt.Println("end")
}