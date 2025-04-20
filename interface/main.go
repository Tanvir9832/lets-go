
package main
import "fmt"

// payment structure
type payment struct {}

// method makePayment for payment
func (p payment) makePayment (amount float64){
	razorpayPaymentGw := razorpay {}
	razorpayPaymentGw.pay(amount)
}

// razorpay structure
type razorpay struct {}

// method pay for razorpay
func (r razorpay) pay (amount float64) {
	fmt.Println("making payment using razorpay", amount)
}

func main(){
	newPayment := payment{}
	newPayment.makePayment(100)
}