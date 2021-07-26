package main

import (
	"fmt"
	"sync"
)
var CurrentBalance int
var mutex = sync.Mutex{}
func deposit(amount int, wg *sync.WaitGroup){
	//mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Println("Request accepted to deposit ",amount)
	CurrentBalance+=amount
	mutex.Unlock()
	wg.Done()
}
func withdraw(amount int, wg *sync.WaitGroup){
	//mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Println("Request accepted to withdraw ", amount)
	if amount>CurrentBalance{
		fmt.Println("Not enough balance.")
	} else{
		CurrentBalance-=amount
	}
	mutex.Unlock()
	wg.Done()
}

func main(){

	CurrentBalance = 500
	wg := sync.WaitGroup{}
	wg.Add(2)
	go deposit(700,&wg)
	go withdraw(1000,&wg)
	wg.Wait()
	fmt.Println("Current balance is: ",CurrentBalance)

}