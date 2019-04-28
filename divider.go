package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	dividerPartLength = 30.0
)

type Divider struct {
	imdraw *imdraw.IMDraw
	start  pixel.Vec
	end    pixel.Vec
}

func getDivider() Divider {
	x := windowLength / 2
	padding := float64(int(windowWidth-barThickness*2) % int(dividerPartLength) / 2)

	divider := Divider{
		start: pixel.V(x, barThickness+padding),
		end:   pixel.V(x, windowWidth-barThickness-padding),
	}

	imdrawPtr := imdraw.New(nil)

	for i, y := 0, divider.start.Y; y < divider.end.Y; i, y = i+1, y+dividerPartLength {
		if i%2 == 0 {
			imdrawPtr.Color = colornames.White
		} else {
			imdrawPtr.Color = colornames.Black
		}

		imdrawPtr.EndShape = endShape
		imdrawPtr.Push(pixel.V(x, y), pixel.V(x, y+dividerPartLength))
	}
	imdrawPtr.Line(barThickness)

	divider.imdraw = imdrawPtr

	return divider
}
