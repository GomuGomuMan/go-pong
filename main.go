package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

const (
	windowLength = 1024.0
	windowWidth  = 798.0
)

func main() {
	pixelgl.Run(run)
}

func getHalfWindowLength() float64 {
	return windowLength / 2
}

func getStaticObjList() []*imdraw.IMDraw {
	var objList []*imdraw.IMDraw

	// Bar
	lowerBarImdraw := GetLowerBar().imdraw
	objList = append(objList, lowerBarImdraw)

	upperBarImdraw := GetUpperBar().imdraw
	objList = append(objList, upperBarImdraw)

	// Divider
	dividerImdraw := getDivider().imdraw
	objList = append(objList, dividerImdraw)

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
	atlas := text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)
	textPadding := 100.0
	textThickness := 30.0
	scoreA := text.New(pixel.V(getHalfWindowLength()-textPadding, windowWidth-80), atlas)
	scoreB := text.New(pixel.V(getHalfWindowLength()+textPadding-textThickness, windowWidth-80), atlas)

	score := 0
	fmt.Fprintf(scoreA, "%d", score)
	fmt.Fprintf(scoreB, "%d", score)

	// Moving bar
	leftBar := GetLeftBar()
	rightBar := GetRightBar()

	for !window.Closed() {
		for _, obj := range staticObjList {
			obj.Draw(window)
		}
		leftBar.imdraw.Draw(window)
		rightBar.imdraw.Draw(window)

		scoreA.Draw(window, pixel.IM.Scaled(scoreA.Orig, 4))
		scoreB.Draw(window, pixel.IM.Scaled(scoreB.Orig, 4))

		window.Update()
	}
}
