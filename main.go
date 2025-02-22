package main

import (
	"ray-tracer-challenge/examples"
	"ray-tracer-challenge/geometry"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	simulation := examples.NewSimulation(
		geometry.NewPoint(0, 1, 0),
		geometry.NewVector(1, 1, 0).Normalize(),
		geometry.NewVector(0, -0.1, 0),
		geometry.NewVector(-0.01, 0, 0))

	simulation.Run()
}
