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
	"sync"
)

// The folder that this file is located in
var carouselDir string

func init() {
        _, caller, _, _ := runtime.Caller(0)
	carouselDir = filepath.Dir(caller)+"/"
}

// Main structure
type Carousel struct {
	engine *qml.Engine
	window *qml.Window
	images map[string]image.Image
	rotateRandomly bool
	animationFile string           // From animations folder
	interval uint                  // The time in milliseconds, to wait before begining animation
	speed uint                     // The time in milliseconds, to complete the animation
	posX uint                      // The position in the main window
	posY uint
	width uint
	height uint
        // Whether to resize the images. If false, bottom and right parts of the
	// images are trimmed
	resizeImages bool
	component qml.Object
	isRunning bool
}

// Initializes and returns a new carousel by provided qml engine and window
func NewCarousel(engine *qml.Engine, win *qml.Window) *Carousel {
	c := &Carousel{animationFile: "slide_to_left"}
	c.engine = engine
	c.window = win
	engine.AddImageProvider("provide", func(path string, width, height int) image.Image {
		re, _ := regexp.Compile(`\d\.\d+$`)
		path = re.ReplaceAllString(path, "")
		img, ok := c.images[path]
		if !ok {
			panic("Image not set for path: "+path)
		}
		return img
	})
	return c
}

func getImage(path string) image.Image {
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
func (c *Carousel) GetEngine() *qml.Engine {
	return c.engine
}
func (c *Carousel) GetWindow() *qml.Window {
	return c.window
}

// Add an image by a relative to the application path. 
// You can add images while the carousel is running too.
func (c *Carousel) AddImage(path string) {
	if c.images == nil {
		c.images = make(map[string]image.Image);
	}
	path, _ = filepath.Abs(path)
	image := getImage(path)
	if c.resizeImages {
		image = resize.Resize(c.width, c.height, image, resize.Lanczos3)
	}
	if c.width == 0 {
		c.width = uint(image.Bounds().Max.X)
	}
	if c.height == 0 {
		c.height = uint(image.Bounds().Max.Y)
	}
	c.images[path] = image
	keys := reflect.ValueOf(c.images).MapKeys()
	paths := make([]string, len(c.images))
	for i, key := range keys {
		paths[i] = key.String()
	}
	c.engine.Context().SetVar("images", strings.Join(paths, "|"))
}

// Set a new list of images to rotate.
// Safe to use while carousel is running.
func (c *Carousel) SetImages(images []string) {
	var wg sync.WaitGroup
	c.images = make(map[string]image.Image);
	for i, name := range images {
		path, _ := filepath.Abs(name)
		images[i] = path
		wg.Add(1)
		go func(name string){
			c.AddImage(name)
			wg.Done()
		}(name)
	}
	wg.Wait()
	imagesStr := strings.Join(images, "|")
	c.engine.Context().SetVar("images", imagesStr)
}
func (c *Carousel) getImage(path string) image.Image {
	return c.images[path]
}
// Set animation qml file from the animations folder
func (c *Carousel) SetAnimationFile(file string) {
	if _, err := os.Stat(carouselDir+"animations/"+file+".qml"); os.IsNotExist(err) {
	    panic("Animation file doesn't exist: "+carouselDir+"animations/"+file+".qml")
	}
	if c.isRunning {
		c.component.Destroy()
		component, err := c.engine.LoadFile(carouselDir+"animations/"+file+".qml")
		obj := component.Create(nil)
		obj.Set("parent", c.window.Root())
		if err != nil {
			panic(err)
		}
		c.component = obj
	}
        c.animationFile = file
}
func (c *Carousel) GetAnimationFile() string {
	return c.animationFile
}

func (c *Carousel) SetInterval(interval uint) {
	c.interval = interval
	c.engine.Context().SetVar("myInterval", int(interval))
}
func (c *Carousel) GetInterval() uint {
	return c.interval
}
func (c *Carousel) SetSpeed(speed uint) {
	c.speed = speed
	c.engine.Context().SetVar("speed", int(speed))
}
func (c *Carousel) GetSpeed() uint {
	return c.speed
}
func (c *Carousel) SetPosX(posX uint) {
	c.posX = posX
	c.engine.Context().SetVar("posX", int(posX))
}
func (c *Carousel) GetPosX() uint {
	return c.posX
}
func (c *Carousel) SetPosY(posY uint) {
	c.posY = posY
	c.engine.Context().SetVar("posY", int(posY))
}
func (c *Carousel) GetPosY() uint {
	return c.posY
}
func (c *Carousel) resizeImagesIfNeeded() {
	var wg sync.WaitGroup
	keys := reflect.ValueOf(c.images).MapKeys()
	if len(keys) > 0 && c.resizeImages {
		key := keys[0].String()
		if c.width == 0 && c.height != 0 {
			c.width = uint(c.images[key].Bounds().Max.X/c.images[key].Bounds().Max.Y)*c.height
		}
		for key, img := range c.images {
			if img.Bounds().Max.Y != int(c.height) || 
			img.Bounds().Max.X != int(c.width) {
				wg.Add(1)
				go func(key string, img image.Image){
					c.images[key] = resize.Resize(c.width, c.height, img, resize.Lanczos3)
					wg.Done()
				}(key, img)
			}
		}
	}
	wg.Wait()
	if c.isRunning {
		c.component.Call("clearImageCache")
	}
}

// Sets width of the carousel. You can put CALCULATE_DIMENSION as argument
// and it will calculate width based on provided height and first image
// to preserve aspect ratio.
func (c *Carousel) SetWidth(width uint) {
	if width == CALCULATE_DIMENSION {
		c.resizeImages = true
	}
	c.width = width
	c.engine.Context().SetVar("myWidth", int(width))
	if c.isRunning {
		c.component.Call("updVars")
	}
}
func (c *Carousel) GetWidth() uint {
	return c.width
}
// Same as SetWidth but for height
func (c *Carousel) SetHeight(height uint) {
	if height == CALCULATE_DIMENSION {
		c.resizeImages = true
	}
	c.height = height
	c.engine.Context().SetVar("myHeight", int(height))
	if c.isRunning {
		c.component.Call("updVars")
	}
}
func (c *Carousel) GetHeight() uint {
	return c.height
}
// Whether carousel should resize the images or crop them
func (c *Carousel) SetResizeImages(resize bool) {
	c.resizeImages = resize
	c.resizeImagesIfNeeded()
}
func (c *Carousel) IsResizingImages() bool {
	return c.resizeImages
}
// Whether the images should rotate randomly. Meaning, when image is shown
// the next images is chosen randomly from one of other images.
func (c *Carousel) SetRotateRandomly(rotateRandomly bool) {
	c.rotateRandomly = rotateRandomly
	c.engine.Context().SetVar("rotateRandomly", rotateRandomly)
}
func (c *Carousel) isRotatingRandomly() bool {
	return c.rotateRandomly;
}
// Shows the carousel and starts rotating.
func (c *Carousel) Run() {
	if len(c.images) == 0 {
		panic("no images added to carousel")
	}
	if c.interval == 0 {
		c.interval = DEFAULT_INTERVAL
	}
	if c.speed == 0 {
		c.speed = DEFAULT_SPEED
	}
	c.resizeImagesIfNeeded()
	c.engine.Context().SetVar("myWidth", int(c.width))
	c.engine.Context().SetVar("myHeight", int(c.height))
	c.engine.Context().SetVar("rotateRandomly", c.rotateRandomly)
	c.engine.Context().SetVar("myInterval", int(c.interval))
	c.engine.Context().SetVar("speed", int(c.speed))
	c.engine.Context().SetVar("posX", int(c.posX))
	c.engine.Context().SetVar("posY", int(c.posY))
	c.isRunning = true
	component, err := c.engine.LoadFile(carouselDir+"animations/"+c.animationFile+".qml")
	if err != nil {
		panic(err)
		return
	}
	obj := component.Create(nil)
	c.component = obj
	obj.Set("parent", c.window.Root())
}

// Removes the carousel
func (c *Carousel) Destroy() {
	c.component.Destroy()
}

// Default rotating speed and interval
const (
        DEFAULT_INTERVAL uint = 1000
        DEFAULT_SPEED = 500
)
const CALCULATE_DIMENSION uint = 0
