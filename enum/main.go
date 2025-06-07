package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

func (s Status) OrderStatus() string {
	switch s {
	case Pending:
		return "Pending"
	case Approved:
		return "Accepted"
	case Rejected:
		return "Rejected"
	default:
		return "Unknown"
	}
}

func main() {
	var s Status = Approved
	fmt.Println(s)
	fmt.Println(s.OrderStatus())
}
