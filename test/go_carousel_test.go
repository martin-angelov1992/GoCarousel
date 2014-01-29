package test

import "github.com/martin-angelov1992/GoCarousel"

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	NewCarousel(engine)
}
