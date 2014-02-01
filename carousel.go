package GoCarousel

import (
	"github.com/niemeyer/qml"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"image/gif"
	"image/jpeg"
	"os"
	"strings"
	"regexp"
	"reflect"
        "runtime"
	"path/filepath"
)

type Carousel struct {
	carouselDir string
	engine *qml.Engine
	window *qml.Window
	images map[string]image.Image
	rotateRandomly bool
	animationFile string
	interval uint
	speed uint
	posX uint
	posY uint
	width uint
	height uint
	resizeImages bool
}

func NewCarousel(engine *qml.Engine, win *qml.Window) *Carousel {
        _, caller, _, _ := runtime.Caller(0)
	c := &Carousel{animationFile: "slide_to_left", interval: 5000, speed: 10}
	c.carouselDir = filepath.Dir(caller)+"/"
	c.engine = engine
	c.window = win
	engine.AddImageProvider("pwd", func(path string, width, height int) image.Image {
		img, _ := c.images[path]
		return img
	})
	return c
}

func (c *Carousel) SetEngine(engine *qml.Engine) {
	
}

func GetImage(path string) image.Image {
	re, _ := regexp.Compile(`^(\w:)?[-\\/._ A-Za-z0-9]+$`)
	if re.MatchString(path) == false {
		panic("invalid file name: "+path)
	}
	parts := strings.Split(path, ".")
	format := parts[len(parts)-1]
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
		case "jpg":
			image, err = jpeg.Decode(f)
		case "png":
			image, err = png.Decode(f)
		default:
			panic("not allowed format: "+format+" of file: "+path)
	}
	if err != nil {
		panic(err)
	}
	return image
}

func (c *Carousel) SetImages(images []string) {
	c.images = make(map[string]image.Image);
	image := GetImage(images[0])
	if c.resizeImages {
		image = resize.Resize(c.width, c.height, image, resize.Lanczos3)
	}
	if c.width == 0 {
		c.width = uint(image.Bounds().Max.X)
	}
	if c.height == 0 {
		c.height = uint(image.Bounds().Max.Y)
	}
	path, _ := filepath.Abs(images[0])

	c.images[path] = image
	for i, name := range images {
		path, _ = filepath.Abs(name)
		images[i] = path
		// Wait only for the first image to finish decoding, the others may be decoded while the slideshow is running
		if i == 0 {
			continue
		}
		go func(){
			image = GetImage(name)
			if c.resizeImages {
				image = resize.Resize(c.width, c.height, image, resize.Lanczos3)
			}
			c.images[path] = image
		}()
	}
	imagesStr := strings.Join(images, "|")
	c.engine.Context().SetVar("images", imagesStr)
}
func (c *Carousel) GetImage(path string) image.Image {
	return c.images[path]
}
func (c *Carousel) SetAnimationFile(file string) {
	if _, err := os.Stat(c.carouselDir+"animations/"+file+".qml"); os.IsNotExist(err) {
	    panic("Animation file doesn't exist: "+c.carouselDir+"animations/"+file+".qml")
	}
        c.animationFile = file
}
func (c *Carousel) SetInterval(interval uint) {
	c.interval = interval
	c.engine.Context().SetVar("myInterval", int(interval))
}

func (c *Carousel) SetSpeed(speed uint) {
	c.speed = speed
	c.engine.Context().SetVar("speed", int(speed))
}

func (c *Carousel) SetPosX(posX uint) {
	c.engine.Context().SetVar("posX", int(posX))
}

func (c *Carousel) SetPosY(posY uint) {
	c.engine.Context().SetVar("posY", int(posY))
}

func (c *Carousel) SetSize(size uint) {
	c.resizeImages = true
	c.height = size
	keys := reflect.ValueOf(c.images).MapKeys()
	if len(keys) > 0 {
		key := keys[0].String()
		c.width = uint(c.images[key].Bounds().Max.X/c.images[key].Bounds().Max.Y)*c.height
		for key, img := range c.images {
			if img.Bounds().Max.Y != int(c.height) || 
			img.Bounds().Max.X != int(c.width) {
				c.images[key] = resize.Resize(c.width, c.height, img, resize.Lanczos3)
			}
		}
	}
}

func (c *Carousel) SetWidth(width uint) {
	c.width = width
}

func (c *Carousel) SetHeight(height uint) {
	c.height = height
}

func (c *Carousel) SetResizeImages(resize bool) {
	c.resizeImages = resize
}

func (c *Carousel) SetRotateRandomly(rotateRandomly bool) {
	c.rotateRandomly = rotateRandomly
	c.engine.Context().SetVar("rotateRandomly", rotateRandomly)
}

func (c *Carousel) Run() {
	if len(c.images) == 0 {
		panic("no images added to carousel")
	}
        _, caller, _, _ := runtime.Caller(0)
	c.engine.Context().SetVar("myWidth", int(c.width))
	c.engine.Context().SetVar("myHeight", int(c.height))
	component, err := c.engine.LoadFile(filepath.Dir(caller)+"/animations/"+c.animationFile+".qml")
	if err != nil {
		panic(err)
		return
	}
	obj := component.Create(nil)
	c.engine.Context().SetVar("rotateRandomly", c.rotateRandomly)
	obj.Set("parent", c.window.Root())
}

func (c *Carousel) Destroy() {

}

func (c *Carousel) Pause() {

}

func (c *Carousel) Hide() {

}
