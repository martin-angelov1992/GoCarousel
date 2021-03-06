package main

import (
	"github.com/niemeyer/qml"
	"github.com/martin-angelov1992/GoCarousel"
	"time"
)

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	base, _ := engine.LoadFile("base.qml")
	win := base.CreateWindow(nil)
	carousel := GoCarousel.NewCarousel(engine, win)
	carousel.SetImages([]string{"../test_images/ubuntu-gopher.png", "../test_images/firefox.png"})
	go func(){
		time.Sleep(3*time.Second)
		carousel.SetAnimationFile("slide_up")
		carousel.SetInterval(3000)
		carousel.AddImage("../test_images/fire.jpg")
		carousel.SetPosX(100)
		carousel.SetPosY(50)
		carousel.SetSpeed(700)
		carousel.SetWidth(300)
		carousel.SetHeight(200)
		carousel.SetResizeImages(true)
		carousel.SetResizeImages(true)
	}()
	win.Show()
	carousel.Run()
	win.Wait()
}
