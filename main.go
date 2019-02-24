package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowLength = 1024.0
	windowWidth  = 798.0
)

func main() {
	pixelgl.Run(run)
}

func getStaticObjList() []*imdraw.IMDraw {
	var objList []*imdraw.IMDraw

	barThickness := 15.0
	lowerBar := imdraw.New(nil)
	lowerBar.Color = colornames.Whitesmoke
	lowerBar.EndShape = imdraw.NoEndShape
	lowerBar.Push(pixel.V(0, barThickness/2), pixel.V(windowLength, barThickness/2))
	lowerBar.Line(barThickness)
	objList = append(objList, lowerBar)

	upperBar := imdraw.New(nil)
	upperBar.Color = colornames.Whitesmoke
	upperBar.EndShape = imdraw.NoEndShape
	upperBar.Push(pixel.V(0, windowWidth-barThickness/2), pixel.V(windowLength, windowWidth-barThickness/2))
	upperBar.Line(barThickness)
	objList = append(objList, upperBar)

	divider := imdraw.New(nil)
	x := windowLength / 2

	dividerPartLength := 30.0
	padding := float64(int(windowWidth-barThickness*2) % int(dividerPartLength) / 2)

	for i, y := 0, barThickness+padding; y < windowWidth-barThickness-padding; i, y = i+1, y+dividerPartLength {
		if i%2 == 0 {
			divider.Color = colornames.White
		} else {
			divider.Color = colornames.Black
		}

		divider.EndShape = imdraw.NoEndShape
		divider.Push(pixel.V(x, y), pixel.V(x, y+dividerPartLength))
	}
	divider.Line(barThickness)
	objList = append(objList, divider)

	return objList
}

func createWindow() *pixelgl.Window {
	log.Println("Creating window")
	config := pixelgl.WindowConfig{
		Title:  "Pong",
		Bounds: pixel.R(0, 0, windowLength, windowWidth),
		VSync:  true,
	}

	window, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	return window
}

func run() {
	window := createWindow()
	staticObjList := getStaticObjList()

	// TODO: Create static textbox with dynamic score

	for !window.Closed() {
		for _, obj := range staticObjList {
			obj.Draw(window)
		}

		window.Update()
	}
}
