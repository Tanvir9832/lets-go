package main

import (
	"fmt"
	"math"
)

//!1

// payment structure
type payment struct{}

// method makePayment for payment
func (p payment) makePayment(amount float64) {
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

// razorpay structure
type razorpay struct{}

// method pay for razorpay
func (r razorpay) pay(amount float64) {
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

//! 1

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float32
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return float64(math.Pi) * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) Perimeter() float64 {
	return 2 * float64(math.Pi) * float64(c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func PrintShape(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func main() {
	newPayment := payment{}
	newPayment.makePayment(100)

	//! Shape
	c := Circle{Radius: 5}
	r := Rectangle{
		Width:  5,
		Height: 4,
	}

	PrintShape(c)
	PrintShape(r)
}
