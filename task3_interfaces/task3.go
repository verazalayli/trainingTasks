package task3_interfaces

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func getFloatFromConsole() float64 {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	flt, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return flt
}

func Task3() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter \"Rectangle\" or \"Circle\" or \"End\"(if you want to finalise program): ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		fmt.Println("You entered:", input)
		if input == "Rectangle" {
			fmt.Print("Enter \"width\": ")
			width := getFloatFromConsole()
			fmt.Print("Enter \"height\": ")
			height := getFloatFromConsole()
			PrintArea(Rectangle{width, height})
		}
		if input == "Circle" {
			fmt.Print("Enter \"radius\": ")
			radius := getFloatFromConsole()
			PrintArea(Circle{radius})
		}
		if input == "End" {
			break
		}
		if input != "Rectangle" && input != "Circle" {
			fmt.Println("Invalid input")
		}
	}
}
