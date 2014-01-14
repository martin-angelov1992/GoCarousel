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
    images := []string{"slide1.jpg", "slide2.jpg", "slide3.jpg"}
    slideshow := new(Slideshow)
    slideshow.setPosition(0, 0) // We want the slideshow to be at the top left of our graphic application
    slideshow.setSize(100)      // Pictures will show with height 100 and width auto
    slideshow.setImages(images) // We specify the image that we want to show
    slideshow.start()           // And now the images begin sliding at the top left
```

License
-------
Apache license
