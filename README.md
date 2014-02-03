GoCarousel
==========
A Go library that allows you to place in your Graphical Application a custom carousel. You can specify: 
* The images to rotate.
* Whether to select the next image in rotation randomly. Current image is excluded from random selection.
* The animation. The available animations are:
  * Rotate right to left
  * Rotate left to right
  * Rotate up to down
  * Rotate down to up
* Interval between switching images.
* Speed of transition.
* Position of the carousel.
* Width and height of the carousel.

Installation on Linux
============
```
$ sudo add-apt-repository ppa:ubuntu-sdk-team/ppa
$ sudo apt-get update
$ sudo apt-get install ubuntu-sdk qtbase5-private-dev qtdeclarative5-private-dev
$ go get github.com/martin-angelov1992/GoCarousel
```
For other operating systems if you manage to install [qml](https://github.com/niemeyer/qml) you should be able to use the library.

Example
=======
```
package main

import (
	"github.com/niemeyer/qml"
	"github.com/martin-angelov1992/GoCarousel"
)

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	base, _ := engine.LoadFile("base.qml")
	win := base.CreateWindow(nil)
	carousel := GoCarousel.NewCarousel(engine, win)
	carousel.SetImages([]string{"../test_images/ubuntu-gopher.png", "../test_images/firefox.png", "../test_images/fire.jpg"})
	carousel.SetWidth(200)
	carousel.SetHeight(GoCarousel.CALCULATE_DIMENSION)
	win.Show()
	carousel.Run()
	win.Wait()
}
```

License
-------
Apache license
