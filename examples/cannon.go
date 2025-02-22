package examples

import (
	"fmt"
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
}

func NewSimulation(point, velocity, gravity, wind geometry.Tuple) Simulation {
	return Simulation{
		Position: point,
		Velocity: velocity,
		environment: environment{
			gravity: gravity,
			wind:    wind,
		},
	}
}

func (s *Simulation) Tick() {
	s.Position = s.Position.Add(s.Velocity)
	s.Velocity = s.Velocity.Add(s.environment.gravity).Add(s.environment.wind)
	fmt.Printf("Pos: %s Velocity: %s\n", s.Position, s.Velocity)
}

func (s *Simulation) Run() {
	for s.Position.Y > 0.0 {
		s.Tick()
	}
}
