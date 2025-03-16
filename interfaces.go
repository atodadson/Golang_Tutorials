package main

import "fmt"

// Implementation of interfaces
type Rectangle struct {
    length float32
    width float32
}

type Square struct {
    length float32
}

type Point struct{
    x float32
    y float32
}

func (r Rectangle) area() float32 {
    return r.length * r.width
}

func (r Rectangle) location() Point {
    return Point{r.length, r.width}
}

func (s Square) area() float32 {
    return s.length * s.length
}

func (s Square) location() Point {
    return Point{x: s.length, y: s.length}
}

type Shape interface {
    area() float32
    location() Point
}

func PrintArea(shape Shape) {
    fmt.Println(shape.area())
}

func PrintLocation(shape Shape) {
    fmt.Println(shape.location())
}

func main() {
    // While struct defines an object, interfaces defines behaviour
    // Declaration of interfaces

//     type Shape interface{
//         area() float32
//         location() Point
//     }

//     myRect := Rectangle{20, 15}
//     fmt.Printf("Area of rectangle with sides %v and %v is: %v \n", myRect.length, myRect.width, myRect.area())
//     fmt.Printf("Rectangle with sides %v and %v has location: %v \n", myRect.length, myRect.width, myRect.location())

    var shape1 Shape = Square{length: 8}
    var shape2 Shape = Rectangle{length:5 ,width:12}

    PrintArea(shape1)
    PrintArea(shape2)

    // Type assertions - Helps to access the underlying fields of the interface
    fmt.Println(shape1.(Square).length) // We accessed length field of Square
    fmt.Println(shape2.(Rectangle).width) // We accessed width field of Rectange



}