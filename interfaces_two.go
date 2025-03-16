package main

import "fmt"

type Point struct {
    x float32
    y float32
}

type Vehicle struct {
    velocity float32
    Point
}

// "*" before the Vehicle reference the instance
func (v *Vehicle) fly() {
    v.velocity = 10
}

func (v *Vehicle) land() {
    v.velocity = 0
}

type Spaceship interface {
    fly()
    land()
    position() Point
}

func main() {
    v := Vehicle{velocity: 0, Point: Point{x: 0,y: 0,},}
    v.fly()
    fmt.Println(v.velocity)
    v.land()
    fmt.Println(v.velocity)
}