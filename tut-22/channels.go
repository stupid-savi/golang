// Channels are used to send data from one gouroutine to another. It is like a pipe

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func processNum(numChan chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("id is", <-numChan)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	chanNum := make(chan int)
// 	wg.Add(1)
// 	go processNum(chanNum, &wg)
// 	chanNum <- 50
// 	wg.Wait()
// }

// Channels are used to send data from one gouroutine to another. It is like a pipe

// sending data from main goroutine to another goroutine
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func processNum(numChan chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range numChan { // with range no need to use arrows
// 		fmt.Println("id is", i)
// 		time.Sleep(time.Second)
// 	}

// }

// func main() {
// 	var wg sync.WaitGroup
// 	chanNum := make(chan int)
// 	wg.Add(1)
// 	go processNum(chanNum, &wg)
// 	for {
// 		chanNum <- rand.Intn(1000)
// 	}
// 	wg.Wait()
// }

// receiving data in main goroutine from another goroutine
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func sendResult(result chan int, num1 int, num2 int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func main() {
// 	var wg sync.WaitGroup
// 	resultChan := make(chan int)
// 	for {
// 		wg.Add(1)
// 		go sendResult(resultChan, rand.Intn(100), rand.Intn(400), &wg)
// 		response := <-resultChan // blocking
// 		time.Sleep(time.Second)
// 		fmt.Println(response)
// 	}

// 	wg.Wait()
// }

// Go routine Synchroniser

// package main

// import (
// 	"fmt"
// )

// func sendResult(boolChan chan bool) {

// 	defer func() {
// 		boolChan <- true
// 	}()

// 	fmt.Println("Processing...")
// }

// func main() {

// 	boolChan := make(chan bool)
// 	go sendResult(boolChan)
// 	<-boolChan // blocking- so alternative of wait group
// use channel blocking if only one goroiutines is available to interact else wait group are more structured feature rich
// }

// Buffer Channels -> Buffer Channels are non blocking
// Creating a bulk email sending queue

package main

import (
	"fmt"
	"time"
)

func sendBulkEmails(emailChan <-chan string, doneChan chan<- bool) {
	defer func() { doneChan <- true }()

	// <-doneChan //  can only send not receove due to doneChan chan<- bool arrow
	// emailChan <- "123@gmail.com" // can only receive not send due to arrow emailChan <-chan string
	for email := range emailChan {
		fmt.Println("Sending Email to :-", email)
		time.Sleep(time.Second)
	}
}

func main() {

	var emailChan = make(chan string, 1000) // by providing size we are making it a buffer channel
	var done = make(chan bool)
	// done <- true --> deadlock because it is blocking
	// <-done
	go sendBulkEmails(emailChan, done)
	for i := 0; i <= 1000; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	// <-emailChan // non blocking till size limit then becomes blocking --> pull the one email from queue
	// <-emailChan // non blocking till size limit then becomes blocking

	fmt.Println("Sending Done")

	// we need to close the buffer channel else after the sending all email it went into deadlock because it will keep receiving for the data in goroutinge

	close(emailChan)

	<-done

}

// multiple channel sending and receiving

// package main

// import "fmt"

// func main() {

// 	chan1 := make(chan int)
// 	chan2 := make(chan string)

// 	go func() {
// 		fmt.Println("Sending from channel1 ")
// 		chan1 <- 123
// 	}()

// 	go func() {
// 		fmt.Println("Sending from channel2 ")
// 		chan2 <- "hello"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case channel1 := <-chan1:
// 			fmt.Println("receiving data", channel1)
// 		case channel2 := <-chan2:
// 			fmt.Println("Receiving data", channel2)
// 		}
// 	}

// }
