package main

import (
	"strings"
	"sync"
)

type Drawable interface {
	Draw()
}

type DrawingObject struct {
	D                 Drawable
	Name              string
	CommunicationType string
	Targets           []string
}

const SQUARE string = "square"
const CIRCLE string = "circle"
const PIPE string = "pipe"

type Square struct {
}

func (s Square) Draw() {
	// TODO: complete the function
}

type Circle struct {
}

func (c Circle) Draw() {
	// TODO: complete the function
}

type Pipe struct {
}

func (p Pipe) Draw() {
	// TODO: complete the function
}

func constructDrawingObject(data []string, ch chan DrawingObject, wg *sync.WaitGroup) {
	// start by verifying which object type to construct that's the first index
	var object = DrawingObject{Name: data[1]}
	switch data[0] {
	case SQUARE:
		object.D = Square{}
	case CIRCLE:
		object.D = Circle{}
	case PIPE:
		object.D = Pipe{}
	default:
		panic("The objects supported are square, circle and pipe")
	}
	if len(data) > 2 {
		// That means it went through the long regex
		// there are two more data in data array
		object.CommunicationType = data[2]
		object.Targets = strings.Split(data[3], ",")
	}

	ch <- object
	wg.Done()
}
