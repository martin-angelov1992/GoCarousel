package GoCarousel

import (
	"github.com/niemeyer/qml"
	"image"
	"image/png"
	"image/gif"
	"image/jpeg"
	"os"
	"strings"
	"regexp"
	"reflect"
)

type Carousel struct {
	engine *qml.Engine
	window *qml.Window
	images map[string]image.Image
	rotateRandomly bool
	animation func()
	interval int
	speed int
	posX int
	posY int
	width int
	height int
}

func arraySearch(needle interface{}, haystack interface{}) int {
        refl := reflect.ValueOf(haystack)
        for i := 0; i < refl.Len(); i++ {
                if refl.Index(i).Interface() == needle {
                        return i
                }
        }
        return -1
}

func LeftToRight() {
}
func RightToLeft() {

}
func UpToDown() {

}
func DownToUp() {

}

func NewCarousel(engine *qml.Engine, win *qml.Window) (error, *Carousel) {
	c := &Carousel{animation: LeftToRight, interval: 5000, speed: 10}
	c.engine = engine
	c.window = win
	engine.AddImageProvider("pwd", func(path string, width, height int) image.Image {
		return c.GetImage(path)
	})
	return nil, c
}

func (c *Carousel) SetEngine(engine *qml.Engine) {
	
}

func GetImage(path string) image.Image {
	allowedFormats := []string{"gif", "jpeg", "png"}
	re, _ := regexp.Compile(`^(\w:)?[-\\/._ A-Za-z0-9]+$`)
	if re.MatchString(path) == false {
		panic("invalid file name: "+path)
	}
	parts := strings.Split(path, ".")
	format := parts[len(parts)-1]
	if arraySearch(format, allowedFormats) == -1 {
		panic("not allowed format: "+format+" of file: "+path)
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var image image.Image
	defer f.Close()
	switch format {
		case "gif":
			image, err = gif.Decode(f)
		case "jpeg":
			image, err = jpeg.Decode(f)
		case "png":
			image, err = png.Decode(f)
	}
	if err != nil {
		panic(err)
	}
	return image
}

func (c *Carousel) SetImages(images []string) {
	c.images = make(map[string]image.Image);
	image := GetImage(images[0])
	if c.width == 0 {
		c.width = image.Bounds().Max.X
	}
	if c.height == 0 {
		c.height = image.Bounds().Max.Y
	}
	c.images[images[0]] = image
	for i, name := range images {
		// Wait only for the first image to finish decoding, the others may be decoded while the slideshow is running
		if i == 0 {
			continue
		}
		go func(){
			c.images[name] = GetImage(name)
		}()
	}
	imagesStr := strings.Join(images, "|")
	c.engine.Context().SetVar("images", imagesStr)
	component, err := c.engine.LoadFile("../carousel.qml")
	if err != nil {
		panic(err)
		return
	}
	obj := component.Create(nil)
	obj.Set("parent", c.window.Root())
}
func (c *Carousel) GetImage(path string) image.Image {
	return c.images[path]
}
func (c *Carousel) SetInterval(interval int) {
	
}

func (c *Carousel) SetPosition(x int, y int) {

}

func (c *Carousel) SetPosX(posX int) {

}

func (c *Carousel) SetPosY(posY int) {
	
}

func (c *Carousel) SetSize(size int) {

}

func (c *Carousel) SetWidth(width int) {

}

func (c *Carousel) SetHeight(height int) {
	
}

func (c *Carousel) SetRotateRandomly(rotateRandomly bool) {
	
}

func (c *Carousel) Run() {
	go func() {
		
	}()
}

func (c *Carousel) Destroy() {

}

func (c *Carousel) Pause() {

}

func (c *Carousel) Hide() {

}

type Animation int

const (
        LEFT_TO_RIGHT Animation = iota
        RIGHT_TO_LEFT
        UP_TO_DOWN
        DOWN_TO_UP
)