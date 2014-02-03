package test

import (
	"github.com/niemeyer/qml"
	"github.com/martin-angelov1992/GoCarousel"
	"testing"
)

var engine  *qml.Engine
var win *qml.Window
var component, base qml.Object
var carousel *GoCarousel.Carousel

func init() {
	qml.Init(nil)
	engine = qml.NewEngine()
	base, _ = engine.LoadFile("../examples/base.qml")	
	win = base.CreateWindow(nil)
	carousel = GoCarousel.NewCarousel(engine, win)
}
// Not much can be tested because from here we do not have access to the qml elements, without extra modification of the qml animation files.
// A better test will be to run the examples in the examples folder and see if everything is shown fine.
// Here a test is considered passed if no method panics.
func TestBasic(t *testing.T) {
	carousel.SetImages([]string{"../test_images/ubuntu-gopher.png", "../test_images/firefox.png", "../test_images/fire.jpg"})
	win.Show()
	carousel.Run()
}
func TestCustom(t *testing.T) {
	carousel.SetImages([]string{"../test_images/ubuntu-gopher.png", "../test_images/firefox.png", "../test_images/fire.jpg"})
	carousel.SetAnimationFile("slide_up")
	carousel.SetInterval(2000)
	carousel.SetRotateRandomly(true)
	carousel.SetPosX(100)
	carousel.SetPosY(50)
	carousel.SetSpeed(300)
	carousel.SetWidth(300)
	carousel.SetHeight(GoCarousel.CALCULATE_DIMENSION)
	win.Show()
	carousel.Run()
}
