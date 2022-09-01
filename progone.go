package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

var option uint8

type Shape interface {
	Perimeter() float32
}

type idealTreangle struct {
	sideTreangle float32
}

func (t idealTreangle) Perimeter() float32 {
	return t.sideTreangle * 3
}

type square struct {
	sideSquare float32
}

func (s square) Perimeter() float32 {
	return s.sideSquare * 4
}

type circle struct {
	radiusCircle float32
}

func (c circle) Perimeter() float32 {
	return c.radiusCircle * 2 * math.Pi
}

func main() {
	err := figure()
	if err != nil {
		fmt.Println(err)
	}

}

func printShapePerimeter(shape Shape) {
	fmt.Println(shape.Perimeter())
}

func figure() error {

	fmt.Println("Select a figure to calculate the perimeter (enter the number):")
	fmt.Println("1) Equilateral triangle")
	fmt.Println("2) Square")
	fmt.Println("3) Circle")
	fmt.Fscan(os.Stdin, &option)

	if option == 1 {
		var inputTreangle float32
		fmt.Println("Enter the length of the side of an equilateral triangle: ")
		fmt.Fscan(os.Stdin, &inputTreangle)
		tr := idealTreangle{inputTreangle}
		fmt.Println("Perimeter of an equilateral triangle: ")
		printShapePerimeter(tr)
		return nil
	} else if option == 2 {
		var inputSquare float32
		fmt.Println("Enter the length of the side of the square: ")
		fmt.Fscan(os.Stdin, &inputSquare)
		sq := square{inputSquare}
		fmt.Println("Perimeter of a square: ")
		printShapePerimeter(sq)
		return nil
	} else if option == 3 {
		var inputCircle float32
		fmt.Println("Enter the radius of the circle: ")
		fmt.Fscan(os.Stdin, &inputCircle)
		ci := circle{inputCircle}
		fmt.Println("Circle perimeter: ")
		printShapePerimeter(ci)
		return nil
	}

	return errors.New("Invalid number")
}
