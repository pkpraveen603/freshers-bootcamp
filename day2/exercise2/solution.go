package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func GetRating(id int,rating chan<-float64, wg *sync.WaitGroup){
	//fmt.Println(id,"started")
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	rating <- float64(rand.Intn(10))
	//fmt.Println(id,"done.")
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	sum := 0.0
	NumStudents := 200.0
	rating := make(chan float64,int(NumStudents))
	for i:=1;i<=int(NumStudents);i++{
		wg.Add(1)
		go GetRating(i,rating,&wg)
	}
	wg.Wait()
	close(rating)
	for value:=range rating{
		sum+=value
	}
	average := sum/NumStudents
	fmt.Println("Average rating is: ",average)
}