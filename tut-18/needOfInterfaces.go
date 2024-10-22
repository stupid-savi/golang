package main

import "fmt"

// convetion for interfaces are that add er at last of the struct for which you are creating the interface
// interface contains only methods
// interfces are like a contract that any struct which containes these methods with exact name and parameters and type will implemets that interface
type paymenter interface {
	makePayment(amount float32, accountNo string)
}

type payment struct {
	// paymentGateway razorpay
	paymentGateway paymenter
}

func (p payment) makePayment(amount float32, accountNo string) {
	// stripeGateway := stripe{}
	// stripeGateway.makePayment(amount, accountNo)
}

// payment method -1
type stripe struct{}

func (s stripe) makePayment(amount float32, accountNo string) {
	fmt.Println("Stripe payment of", amount, "INR is done from account no.", accountNo)
}

// type payment method -2
type razorpay struct{}

func (r razorpay) makePayment(amount float32, accountNo string) {
	fmt.Println("Razorpay payment of", amount, "INR is done from account no.", accountNo)
}

// type payment method -3

type fakePayment struct{}

func (f fakePayment) makePayment(amount float32, accountNo string) {
	fmt.Println("Testing payment of", amount, "INR is done from account no.", accountNo)
}

func main() {
	// gatway
	// razorPay := razorpay{}
	// stripe := stripe{}
	fakePayment := fakePayment{}
	myPayment := payment{
		// paymentGateway: stripe,
		// paymentGateway: razorPay,
		paymentGateway: fakePayment,
	}
	myPayment.paymentGateway.makePayment(123, "99090909090")

}

// 1. The problem is that if tomorrow I need to add second interface then I nned to change the hard codded payment method in our makePayment method So we will remove the method and struct initialization from the makePayment method and add razorpya to it

//2. Still the paymentgateway is hard coded and it voilates the open close principle which is open for extenstion and close for modifcation type payment struct {
//paymentGateway razorpay
//}

// Because adding a gatway for testing and paypal we need to change the hard coded value of the payment gatway everytime

// Let's make a interface to solve this problem
// If we add another method to paymenter tthen need to add it to all struct using it as payment gateway
