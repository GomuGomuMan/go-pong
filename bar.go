package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	barThickness    = 15.0
	endShape        = imdraw.NoEndShape
	playerBarLength = 100
)

type Bar struct {
	imdraw *imdraw.IMDraw
	start  pixel.Vec
	end    pixel.Vec
	length float64
}

func GetLowerBar() Bar {
	lowerBar := Bar{
		start:  pixel.V(0, barThickness/2),
		end:    pixel.V(windowLength, barThickness/2),
		length: windowWidth,
	}

	imdrawPtr := imdraw.New(nil)
	imdrawPtr.Color = colornames.Whitesmoke
	imdrawPtr.EndShape = endShape
	imdrawPtr.Push(lowerBar.start, lowerBar.end)
	imdrawPtr.Line(barThickness)

	lowerBar.imdraw = imdrawPtr

	return lowerBar
}

func GetUpperBar() Bar {
	upperBar := Bar{
		start:  pixel.V(0, windowWidth-barThickness/2),
		end:    pixel.V(windowLength, windowWidth-barThickness/2),
		length: windowWidth,
	}

	imdrawPtr := imdraw.New(nil)
	imdrawPtr.Color = colornames.Whitesmoke
	imdrawPtr.EndShape = imdraw.NoEndShape
	imdrawPtr.Push(upperBar.start, upperBar.end)
	imdrawPtr.Line(barThickness)

	upperBar.imdraw = imdrawPtr

	return upperBar
}

func GetLeftBar() Bar {
	length := 100.0
	leftBar := Bar{
		start:  pixel.V(barThickness/2, windowWidth/2-length/2),
		end:    pixel.V(barThickness/2, windowWidth/2+length/2),
		length: length,
	}

	imdrawPtr := imdraw.New(nil)
	imdrawPtr.Color = colornames.Whitesmoke
	imdrawPtr.EndShape = imdraw.NoEndShape
	imdrawPtr.Push(leftBar.start, leftBar.end)
	imdrawPtr.Line(barThickness)

	leftBar.imdraw = imdrawPtr

	return leftBar
}

func GetRightBar() Bar {
	length := 100.0
	rightBar := Bar{
		start:  pixel.V(windowLength-barThickness/2, windowWidth/2-length/2),
		end:    pixel.V(windowLength-barThickness/2, windowWidth/2+length/2),
		length: length,
	}

	imdrawPtr := imdraw.New(nil)
	imdrawPtr.Color = colornames.Whitesmoke
	imdrawPtr.EndShape = imdraw.NoEndShape
	imdrawPtr.Push(rightBar.start, rightBar.end)
	imdrawPtr.Line(barThickness)

	rightBar.imdraw = imdrawPtr

	return rightBar
}
