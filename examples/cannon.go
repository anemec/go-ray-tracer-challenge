package examples

import (
	"fmt"
	"math"
	"ray-tracer-challenge/canvas"
	"ray-tracer-challenge/geometry"
)

type projectile struct {
	point    geometry.Tuple
	velocity geometry.Tuple
}

type environment struct {
	gravity geometry.Tuple
	wind    geometry.Tuple
}

type Simulation struct {
	Position    geometry.Tuple
	Velocity    geometry.Tuple
	environment environment
	Image       canvas.Canvas
}

func NewSimulation(point, velocity, gravity, wind geometry.Tuple, canvas canvas.Canvas) Simulation {
	yPos := canvas.Height - int(math.Round(point.Y))
	xPos := int(math.Round(point.X))
	canvas.WritePixel(yPos, xPos, geometry.NewColor(0, 1, 0))
	return Simulation{
		Position: point,
		Velocity: velocity,
		environment: environment{
			gravity: gravity,
			wind:    wind,
		},
		Image: canvas,
	}
}

// Tick TODO cleanup boundary checks
func (s *Simulation) Tick() {
	s.Position = s.Position.Add(s.Velocity)
	s.Velocity = s.Velocity.Add(s.environment.gravity).Add(s.environment.wind)
	yPos := s.Image.Height - int(math.Round(s.Position.Y))
	xPos := int(math.Round(s.Position.X))
	if (yPos >= 0 && yPos <= s.Image.Height) && (xPos >= 0 && xPos <= s.Image.Width) {
		s.Image.WritePixel(yPos, xPos, geometry.NewColor(1, 0, 0))
	}
	fmt.Printf("Pos: %s Velocity: %s\n", s.Position, s.Velocity)
}

func (s *Simulation) Run() {
	for s.Position.Y > 0.0 {
		s.Tick()
	}
}
