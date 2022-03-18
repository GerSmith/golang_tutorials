/*
	Программирование на Golang. Структуры
*/

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

type Vector struct {
	x float64
	y float64
	z float64
}

func createVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z}
}

func length(v Vector) int {
	return int(math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z))
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2.0)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Printf("x = %.2f, y = %.2f, r = %.2f\n", c.x, c.y, c.r)
	fmt.Printf("Circle area = %.2f\n", c.area())

	d := Rectangle{0, 0, 10, 10}
	fmt.Printf("x1 = %.2f, y1 = %.2f, x2 = %.2f, y2 = %.2f\n", d.x1, d.y1, d.x2, d.y2)
	fmt.Printf("Rectangle area = %.2f\n", d.area())

	human := Person{"Artem"}
	fmt.Printf("Hello %v! Lets talk!\n", human.Name)
	human.Talk()

	and := new(Android)
	and.Person.Talk()

	var C3PO = Android{
		Model: "Droid",
		Person: Person{
			Name: "C3PO",
		},
	}
	fmt.Printf("%v - %v\n", C3PO.Model, C3PO.Name)
	C3PO.Person.Talk()

	a := createVector(6, 3, 2)
	b := createVector(1, 2, 4)
	fmt.Print(length(a), " ")
	fmt.Print(length(b))
}
