package main

import (
	"fmt"
)
func UpdateCount(id int,message <-chan string, sendMap chan<- map[string] int) {

	str := <-message
	freq := make(map[string]int)
	fmt.Println("string",id,"calculation started.")
	for i:=0;i<len(str);i++{
		freq[str[i:i+1]]+=1
	}
	sendMap<-freq
	fmt.Println("String",id,"calculation done.")
}

func main(){

	GivenString := [] string{"quick","brown","fox","lazy","dog"}
	message := make(chan string,len(GivenString))
	recieveMap := make(chan map[string]int, len(GivenString) )
	for i:=0;i<len(GivenString);i++{
		go UpdateCount(i,message,recieveMap)
	}
	for i:=0;i<len(GivenString);i++{
		message <- GivenString[i]
	}
	counter := make(map[string]int)
	for i:=0;i<len(GivenString);i++{
		tempCounter:=<-recieveMap
		for k,v := range tempCounter{
			counter[k]+=v
		}
	}
	//time.Sleep(2*time.Second)
	for k,v:=range counter{
		fmt.Println(k,v)
	}
}