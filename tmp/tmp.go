package main

import "fmt"

func main(){
	for x := 0 ; x<=10;x+=1{
		fmt.Println("Leo:",x)
		if x==2{
			fmt.Println("ola")
		}
	}
}